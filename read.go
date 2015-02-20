package main

import (
	"gopkg.in/yaml.v2"
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
		if file == nil {
			return nil
		}

		err = ReadFile(file)
		if err != nil {
			return err
		}
		files = append(files, file)
		return nil
	})

	return
}

func ReadFile(file *File) (err error) {
	if file.IsReadContent() == false {
		return
	}

	var c []byte
	c, err = ioutil.ReadFile(file.Path)
	if err != nil {
		return err
	}

	if file.Type == FileTypeMarkdown {

		header, body, _err := ReadMetaAndBody(c)
		if _err != nil {
			return _err
		}
		file.Meta = header
		file.Content = body

	} else if file.Type == FileTypeYAMLConf {

		var meta map[string]interface{}
		if err = yaml.Unmarshal(c, &meta); err != nil {
			return
		}

		file.Meta = meta

	} else {

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
			if err = yaml.Unmarshal(header, &meta); err != nil {
				return
			}
		}
	} else {
		body = content
	}

	return
}
