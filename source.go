package main

import (
	"log"
	"os"
	"path/filepath"
)

type Source struct {
	path    string
	content string
}

type Sources []Source

func NewSources(srcPath string) (Sources, error) {
	srcs := Sources{}

	filepath.Walk(srcPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Println(err)
			return err
		}

		if isFile(path) == false {
			return nil
		}

		var src Source
		src, err = ReadSource(path)
		if err != nil {
			return err
		}
		srcs = append(srcs, src)
		return nil
	})

	return srcs, nil
}

func isFile(path string) (isFile bool) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		isFile = false
		return
	}

	if fileInfo.IsDir() {
		isFile = false
	} else {
		isFile = true
	}
	return

}
