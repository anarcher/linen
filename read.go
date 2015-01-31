package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func Read(srcs Sources, path string) (err error) {

	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		src := Source{}
		err = ReadOne(src, path)
		srcs[Path(path)] = src

		return err
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

	default:
		src["type"] = SourceTypeCopy

	}

	return

}
