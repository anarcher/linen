package main

import (
	"github.com/mgutz/logxi/v1"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type previewHandler struct {
	files  Files
	build  bool
	logger log.Logger
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

	p.logger.Debug("file", "path", path)

	files, err := ReadFiles(SrcPath)
	if err != nil {
		logger.Error("Read", "err", err)
		return
	}

	err = TransformFiles(files)
	if err != nil {
		logger.Error("Transform", "err", err)
	}

	if p.build {
		err = WriteFiles(files, TargetPath)
		if err != nil {
			logger.Error("Write", "err", err)
		}
	}

	for _, file := range files {

		if p.logger.IsDebug() {
			p.logger.Debug("file", "path", path, "url", file.Url())
		}

		if file.Url() == path {
			//Render file
			_, err = w.Write(file.Content)
			if err != nil {
				p.logger.Error("response", "err", err)
			}
			return
		}
	}

	http.NotFound(w, r)
	return
}

func previewServe(addr string, build bool) {

	p := &previewHandler{
		build:  build,
		logger: log.New("preview")}

	http.Handle("/", p)

	p.logger.Info("serve", "addr", addr, "src", SrcPath, "target", TargetPath, "build", build)

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Error("preview", "err", err)
		os.Exit(1)
	}
}
