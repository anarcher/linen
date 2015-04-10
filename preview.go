package main

import (
	"fmt"
	"github.com/mgutz/logxi/v1"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type previewHandler struct {
	srcPath, targetPath string
	files               Files
	build               bool
	logger              log.Logger
}

func (p *previewHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	if strings.HasSuffix(path, "/") {
		path = path + "index.html"
	} else if filepath.Ext(path) == "" {
		path = path + "/index.html"
	}

	ext := filepath.Ext(path)

	if ext == ".html" {
		path1 := fmt.Sprintf("%s%s", p.srcPath, path)
		p.logger.Debug("path", "path", path)
		if _, err := os.Stat(path1); err != nil {

			path1 = fmt.Sprintf("%s%s%s", p.srcPath, strings.TrimRight(path, ".html"), ".md")
			p.logger.Debug("path1", "path", path1)

			if _, err := os.Stat(path1); err != nil {
				p.logger.Warn("NotFound", "path", path)
				http.NotFound(w, r)
				return
			}
		}
		path = path1
	}

	p.logger.Debug("path", "path", path)

	files := p.files
	var file *File
	fileIdx := -1

	for i, f := range files {
		if f.Path() == path {
			file = f
			fileIdx = i
			break
		}
	}
	if file == nil {
		fi, err := os.Stat(path)
		if err != nil {
			p.logger.Error("file", "err", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		file = NewFile(path, fi)
	}

	err := ReadFile(file)
	if err != nil {
		p.logger.Error("read", "err", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if fileIdx == 0 {
		files = append(files[1:], file)
	} else if fileIdx > 0 {
		files = append(files[0:fileIdx-1], file)
		files = append(files, files[fileIdx:]...)
	} else if fileIdx == -1 {
		files = append(files, file)
	}

	//Transform
	err = TransformFile(file, files)
	if err != nil {
		p.logger.Error("transform", "err", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//Build
	if p.build == true {
		err = WriteFile(file, path)
		if err != nil {
			p.logger.Error("transform", "err", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	//Render file
	_, err = w.Write(file.Content)
	if err != nil {
		p.logger.Error("response", "err", err)
	}

	return
}

func previewServe(addr, srcPath, targetPath string, build bool) {

	files, err := ReadFiles(srcPath)
	if err != nil {
		logger.Error("Read", "err", err)
		return
	}

	err = TransformFiles(files)
	if err != nil {
		logger.Error("Transform", "err", err)
	}

	if build {
		err = WriteFiles(files, targetPath)
		if err != nil {
			logger.Error("Write", "err", err)
		}
	}
	p := &previewHandler{
		files:      files,
		srcPath:    srcPath,
		targetPath: targetPath,
		build:      build,
		logger:     log.New("preview")}

	http.Handle("/", p)

	p.logger.Info("serve", "addr", addr, "src", srcPath, "target", targetPath, "build", build)

	err = http.ListenAndServe(addr, nil)
	if err != nil {
		log.Error("preview", "err", err)
		os.Exit(1)
	}
}
