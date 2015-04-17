package main

import (
	"github.com/ahmetalpbalkan/go-linq"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"time"
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

func (fs Files) Filter(args ...interface{}) linq.Query {
	var query linq.Query
	queryOrT := args[len(args)-1]
	if _, ok := queryOrT.(linq.Query); ok {
		query = queryOrT.(linq.Query)
	} else {
		query = linq.From(fs)
	}

	var _args []string
	for _, a := range args[:len(args)-1] {
		_args = append(_args, a.(string))
	}
	exprs, err := parseExprs(_args)
	if err != nil {
		panic(err)
	}

	for _, expr := range exprs {
		whereFunc := expr.WhereFunc()
		query = query.Where(whereFunc)
	}

	return query
}

func (fs Files) Sort(args ...interface{}) linq.Query {
	query := args[len(args)-1].(linq.Query)
	var _args []string
	for _, a := range args[:len(args)-1] {
		_args = append(_args, a.(string))
	}

	exprs, err := parseExprs(_args)
	if err != nil {
		panic(err)
	}

	for _, expr := range exprs {
		orderByFunc := expr.OrderByFunc()
		query = query.OrderBy(orderByFunc)
	}

	return query
}

func (fs Files) Group(key string, query linq.Query) map[linq.T][]linq.T {
	groupByFunc := func(t linq.T) linq.T {
		file := *(t.(*File))
		keyValue := reflect.ValueOf(file).FieldByName(key)
		return keyValue
	}
	q, err := query.GroupBy(groupByFunc, groupByFunc)
	if err != nil {
		panic(err)
	}
	return q
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

	dir := strings.Replace(filepath.Dir(path), SrcPath, "", 1)
	if dir == "" {
		dir = "/"
	}
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
	return filepath.Join(SrcPath, f.Dir, f.Base)
}

func (f File) Url() string {
	var fileName string
	if f.Ext == ".md" {
		fileName = filepath.Join("/", f.Dir, strings.Replace(f.Base, f.Ext, ".html", 1))

	} else {
		fileName = filepath.Join("/", f.Dir, f.Base)
	}
	return fileName
}

func (f File) Date() (time.Time, error) {
	const layout = "2006-01-01"
	if date, ok := f.Meta["date"]; ok {
		t, err := time.Parse(layout, date.(string))
		return t, err
	}

	return f.Info.ModTime(), nil
}
