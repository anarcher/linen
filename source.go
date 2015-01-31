package main

import (
	"github.com/stretchr/powerwalk"
	"io/ioutil"
	"log"
	"os"
)

type Source struct {
	path    string
	content string
}

type Sources []Source

func (s Sources) Write(destPath string) error {
	for i, src := range s {
		log.Println(i, src)
	}
	return nil
}

func NewSources(srcPath string) (Sources, error) {
	srcs := Sources{}

	powerwalk.Walk(srcPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Println(err)
			return err
		}

		if info.IsDir() == true {
			return nil
		}

		var content []byte

		content, err = ioutil.ReadFile(path)
		if err != nil {
			log.Println(err)
			return err
		}

		src := Source{path: path, content: string(content)}
		srcs = append(srcs, src)
		return nil
	})

	return srcs, nil
}
