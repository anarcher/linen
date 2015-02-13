package main

import (
	"encoding/json"
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

		file := NewFile(path, info)
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
		header, body, _err := ReadMetaAndBody(c)
		if _err != nil {
			return _err
		}
		file.Meta = header
		file.Content = body

	} else {
		var c []byte
		c, err = ioutil.ReadFile(file.Path)
		if err != nil {
			return err
		}
		file.Content = c
	}

	return nil
}

func ReadMetaAndBody(content []byte) (meta map[string]interface{}, body []byte, err error) {

	c := string(content)
	if len(c) <= 0 {
		return
	}

	cs := strings.SplitN(c, "\n---\n", 2)

	if len(cs) == 2 {
		body = []byte(cs[1])
		header := []byte(cs[0])
		if len(header) > 0 {
			if err = json.Unmarshal(header, &meta); err != nil {
				return
			}
		}
	} else {
		body = content
	}

	return
}
