package main

import (
	"os"
	"path/filepath"
)

const (
	FileTypeMarkdown = iota
	FileTypeYAMLConf
	FileTypeTemplate
	FileTypeHTML
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

	if base := filepath.Base(path); base == "_.yaml" {
		fileType = FileTypeYAMLConf
	}

	file.Type = fileType
	return file
}

func (f *File) IsWrite() bool {
	if f.Type == FileTypeYAMLConf || f.Type == FileTypeTemplate {
		return false
	}
	return true
}

func (f *File) IsReadContent() bool {
	if f.Type == FileTypeTemplate {
		return false
	}
	return true
}
