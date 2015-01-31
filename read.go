package main

import (
	"io/ioutil"
	"path/filepath"
)

func ReadSource(path string) (source Source, err error) {
	source.Path = path
	source.Ext = filepath.Ext(path)

	switch source.Ext {
	case ".md":
		source.Type = SourceTypeMarkdown
		err = ReadSourceContent(&source)
	case ".template":
		source.Type = SourceTypeTemplate
		err = ReadSourceContent(&source)
	default:
		source.Type = SourceTypeCopy
	}

	return
}

func ReadSourceContent(src *Source) (err error) {
	var content []byte
	content, err = ioutil.ReadFile(src.Path)
	if err != nil {
		return
	}

	src.Content = string(content)
	return

}

func ReadSourceHeader(source *Source) (err error) {
	return nil
}
