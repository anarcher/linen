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
		file.Content["header"] = header
		file.Content["body"] = body

	} else if file.Type == FileTypeTemplate {
		var c []byte
		c, err = ioutil.ReadFile(file.Path)
		if err != nil {
			return err
		}
		file.Content["body"] = string(c)
	}

	return nil
}

func ReadHeaderAndBody(content []byte) (header, body string, err error) {

	c := string(content)
	if len(c) <= 0 {
		return
	}

	cs := strings.SplitN(c, "\n---\n", 2)

	if len(cs) == 2 {
		header = cs[0]
		body = cs[1]
	} else {
		body = c
	}

	return
}
