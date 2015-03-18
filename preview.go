package main

import (
	"github.com/mgutz/logxi/v1"
	"net/http"
	"os"
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
			p.logger.Error("file", "err", err)
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
	files, err := Build(srcPath, targetPath)

	http.Handle("/", &previewHandler{
		files:      files,
		srcPath:    srcPath,
		targetPath: targetPath,
		build:      build,
		logger:     log.New("preview")})

	err = http.ListenAndServe(addr, nil)
	if err != nil {
		log.Error("preview", "err", err)
		os.Exit(1)
	}
}
