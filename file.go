package main

import (
	"github.com/ahmetalpbalkan/go-linq"
	"os"
	"path/filepath"
	"strings"
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
	Dir     string
	Base    string
	Ext     string
	Type    int
	Meta    FileMeta
	Info    os.FileInfo
	Content []byte
}

type Files []*File

func (fs Files) Filter(queryOrT interface{}, args ...string) linq.Query {
	var query linq.Query
	if _, ok := queryOrT.(linq.Query); ok {
		query = queryOrT.(linq.Query)
	} else {
		query = linq.From(fs)
	}

	exprs, err := parseExprs(args)
	if err != nil {
		panic(err)
	}

	for _, expr := range exprs {
		whereFunc := expr.WhereFunc()
		query = query.Where(whereFunc)
	}

	return query
}

func (fs Files) Sort(query linq.Query, args ...string) linq.Query {
	exprs, err := parseExprs(args)
	if err != nil {
		panic(err)
	}

	for _, expr := range exprs {
		orderByFunc := expr.OrderByFunc()
		query = query.OrderBy(orderByFunc)
	}

	return query
}

func (fs Files) Results(query linq.Query) []linq.T {
	results, err := query.Results()
	if err != nil {
		panic(err)
	}
	return results
}

func (fs Files) Count(query linq.Query) int {
	cnt, err := query.Count()
	if err != nil {
		panic(err)
	}
	return cnt
}

func NewFile(path string, info os.FileInfo) *File {

	dir := filepath.Dir(path)
	base := filepath.Base(path)
	ext := filepath.Ext(path)

	file := &File{Dir: dir, Base: base, Ext: ext, Info: info, Meta: make(map[string]interface{})}

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

	if base == "_.yaml" {
		fileType = FileTypeYAMLConf
	}

	file.Type = fileType
	return file
}

func (f File) IsWrite() bool {
	if f.Type == FileTypeYAMLConf || f.Type == FileTypeTemplate {
		return false
	}
	return true
}

func (f File) IsReadContent() bool {
	if f.Type == FileTypeTemplate {
		return false
	}
	return true
}

func (f File) Path() string {
	return f.Dir + string(os.PathSeparator) + f.Base
}

func (f File) Url() string {
	var fileName string
	if f.Ext == ".md" {
		fileName = f.Dir + strings.Replace(f.Base, f.Ext, ".html", 1)

	} else {
		fileName = f.Path()
	}
	return fileName
}
