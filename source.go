package main

import (
	"log"
	"os"
	"path/filepath"
)

const (
	SourceTypeMarkdown SourceType = iota
	SourceTypeTemplate
	SourceTypeCopy
)

type SourceType int

type Source struct {
	Type    SourceType
	Header  map[string]interface{}
	Path    string
	Content string
	Body    string
	Ext     string
}

type Sources []Source

func NewSources(srcPath string) (Sources, error) {
	srcs := Sources{}

	filepath.Walk(srcPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Println(err)
			return err
		}

		if info.IsDir() == true {
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
