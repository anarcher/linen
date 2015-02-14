package main

import (
	"os"
	"path/filepath"
)

const (
	FileTypeMarkdown = iota
	FileTypeTemplate
	FileTypeHTML
	FileTypeJSON
	FileTypePlain
)

type FileMeta map[string]interface{}

type File struct {
	Path    string
	Ext     string
	Type    int
	Meta    FileMeta
	Info    os.FileInfo
	Content []byte
}

type Files []*File

func NewFile(path string, info os.FileInfo) *File {
	ext := filepath.Ext(path)
	file := &File{Path: path, Ext: ext, Info: info, Meta: make(map[string]interface{})}
	var fileType int
	switch ext {
	case ".md":
		fileType = FileTypeMarkdown
	case ".tmpl":
		fileType = FileTypeTemplate
	case ".html":
		fileType = FileTypeHTML
	default:
		fileType = FileTypePlain
	}
	file.Type = fileType
	return file
}
