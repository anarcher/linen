package main

import (
	"fmt"
	"io/ioutil"
	"os"
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
	destPath := path + file.Path
	fmt.Println("destPath:" + destPath)
	os.MkdirAll(destPath, DirPerm)

	err = ioutil.WriteFile(destPath, file.Content, file.Info.Mode())

	return
}
