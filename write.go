package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const (
	DirPerm = 0774
)

func WriteFiles(files Files, path string) (err error) {
	for _, file := range files {
		err = WriteFile(file, path)
		if err != nil {
			return
		}
	}
	return
}

func WriteFile(file *File, path string) (err error) {
	dirPath := path + "/" + filepath.Dir(file.Path)
	os.MkdirAll(dirPath, DirPerm)

	fullPath := filepath.Join(path, "/")
	if file.Type == FileTypeMarkdown {
		fullPath = filepath.Join(fullPath, filepath.Dir(file.Path), "/", strings.Replace(filepath.Base(file.Path), file.Ext, ".html", 1))
	} else {
		fullPath = filepath.Join(fullPath, file.Path)
	}

	if file.Type != FileTypeTemplate {
		err = ioutil.WriteFile(fullPath, file.Content, file.Info.Mode())
	}

	return
}
