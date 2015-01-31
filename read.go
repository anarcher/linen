package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func Read(srcs Sources, path string) (err error) {

	err = filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		src := Source{}
		err = ReadOne(src, path)
		if err != nil {
			return err
		}
		srcs[Path(path)] = src
		return nil
	})

	return
}

func ReadOne(src Source, path string) (err error) {

	fileExt := filepath.Ext(path)
	src["ext"] = fileExt

	switch fileExt {
	case ".md":
		src["type"] = SourceTypeMarkdown

		var c []byte
		c, err = ioutil.ReadFile(path)
		if err != nil {
			return
		}

		src["content"] = c
		err = readHeaderAndBody(src)

	case ".tmpl":
		src["type"] = SourceTypeTemplate
	default:
		src["type"] = SourceTypeCopy
	}

	return

}

func readHeaderAndBody(src Source) (err error) {
	c := string(src["content"].([]byte))
	if len(c) <= 0 {
		return
	}

	cs := strings.SplitN(c, "\n---\n", 2)

	if len(cs) == 2 {
		var data map[string]interface{}
		err = json.Unmarshal([]byte(cs[0]), &data)
		if err != nil {
			return
		}
		src["header"] = data
		src["body"] = cs[1]
	} else {
		src["body"] = c
	}
	return
}
