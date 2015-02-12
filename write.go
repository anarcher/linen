package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
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
	fullPath := path + "/" + file.Path
	os.MkdirAll(dirPath, DirPerm)

	err = ioutil.WriteFile(fullPath, file.Content, file.Info.Mode())

	return
}
