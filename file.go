package main

import (
	"path/filepath"
)

const (
	FileTypeMarkdown = iota
	FileTypeTemplate
	FileTypePlain
)

const (
	FileContentHeader = "header"
	FileContentBody   = "body"
)

type FileContent map[string]interface{}

type File struct {
	Path    string
	Ext     string
	Type    int
	Content FileContent
}

type Files []*File

func NewFile(path string) *File {
	ext := filepath.Ext(path)
	file := &File{Path: path, Ext: ext}
	var fileType int
	switch ext {
	case ".md":
		fileType = FileTypeMarkdown
		file.Content = make(map[string]interface{})
	case ".tmpl":
		fileType = FileTypeTemplate
		file.Content = make(map[string]interface{})
	default:
		fileType = FileTypePlain
	}
	file.Type = fileType
	return file
}
