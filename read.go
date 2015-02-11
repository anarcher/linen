package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func ReadFiles(path string) (files Files, err error) {

	err = filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return err
		}

		file := NewFile(path)
		ReadFileContent(file)
		files = append(files, file)
		return nil
	})

	return
}

func ReadFileContent(file *File) (err error) {
	if file.Type == FileTypeMarkdown {
		var c []byte
		c, err = ioutil.ReadFile(file.Path)
		if err != nil {
			return err
		}
		header, body, _err := ReadHeaderAndBody(c)
		if _err != nil {
			return _err
		}
		file.Content[FileContentHeader] = header
		file.Content[FileContentBody] = body

	} else if file.Type == FileTypeTemplate {
		var c []byte
		c, err = ioutil.ReadFile(file.Path)
		if err != nil {
			return err
		}
		file.Content[FileContentBody] = c
	}

	return nil
}

func ReadHeaderAndBody(content []byte) (header, body []byte, err error) {

	c := string(content)
	if len(c) <= 0 {
		return
	}

	cs := strings.SplitN(c, "\n---\n", 2)

	if len(cs) == 2 {
		header = []byte(cs[0])
		body = []byte(cs[1])
	} else {
		body = content
	}

	return
}
