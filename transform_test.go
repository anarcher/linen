package main

import (
	"github.com/kr/pretty"
	"html/template"
	"os"
	"testing"
)

func TestTransformFiles(t *testing.T) {
	path := "./examples/basic"
	files, err := ReadFiles(path)
	if err != nil {
		t.Error(err)
	}

	err = TransformFiles(files)
	if err != nil {
		t.Error(err)
	}

	t.Logf("%# v", pretty.Formatter(files))

}

func TestTransformFilesWithTemplateMeta(t *testing.T) {
	path := "./examples/layout"
	files, err := ReadFiles(path)
	if err != nil {
		t.Error(err)
	}

	err = TransformFiles(files)
	if err != nil {
		t.Error(err)
	}

	t.Logf("len:%d", len(files))

	//t.Logf("%# v", pretty.Formatter(files))

	for i, file := range files {
		t.Logf("%d,%s", i, file.Content)
	}
}

func TestFileTemplate(t *testing.T) {
	path := "./examples/layout/article01.md"
	info, err := os.Stat(path)
	if err != nil {
		t.Error(err)
	}

	t.Logf("%v", info)
	file := NewFile(path, info)
	err = ReadFile(file)
	if err != nil {
		t.Error(err)
	}

	var tmpl *template.Template
	tmpl, err = transformTemplateMeta(file)
	if err != nil {
		t.Error(err)
	}

	t.Logf("file.Content:%s", file.Content)
	t.Logf("file.Meta:%v", file.Meta)
	templateFile, exists := file.Meta["template_file"]

	t.Logf("template_file.exists:%s", exists)
	if exists == true {
		t.Logf("%v %s", templateFile.(string), exists)
	}

	if tmpl == nil {
		t.Error("tmpl is nil!!!")
	}

	t.Logf("%v", tmpl)

}
