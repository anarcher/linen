package main

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/russross/blackfriday"
	"html/template"
	"os"
	"path/filepath"
	//	"github.com/ahmetalpbalkan/go-linq"
)

var (
	TemplateFileNotFound = errors.New("TemplateFile not found")
)

func TransformFiles(files Files) (err error) {
	for _, file := range files {
		err = TransformFile(file)
		if err != nil {
			return
		}
	}
	return
}

func TransformFile(file *File) (err error) {
	if file.Type == FileTypeMarkdown {
		file.Content = blackfriday.MarkdownBasic(file.Content)
	}

	if file.Type == FileTypeMarkdown ||
		file.Type == FileTypeHTML {

		var tmpl *template.Template

		if file.Type == FileTypeMarkdown {
			tmpl, err = transformTemplateMeta(file)
		} else {
			tmpl, err = transformTemplate(file)
		}

		//TODO: More expressive about transformTemplate's works
		if err != nil || tmpl == nil {
			return
		}

		var output bytes.Buffer

		err = tmpl.Execute(&output, file.Meta)
		if err != nil {
			return err
		}
		file.Content = output.Bytes()
	}
	return
}

func transformTemplateMeta(file *File) (tmpl *template.Template, err error) {
	tf, exists := file.Meta["template_file"]
	if exists == false {
		return
	}

	templateFile, ok := tf.(string)
	if !ok {
		err = errors.New("template_file should be string type.")
		return
	}

	if filepath.IsAbs(templateFile) == false {
		fileDirPath := filepath.Dir(file.Path)
		templateFile, err = filepath.Abs(filepath.Join(fileDirPath, templateFile))
		if err != nil {
			return
		}
	}

	if _, err = os.Stat(templateFile); err != nil {
		err = TemplateFileNotFound
		return
	}

	content := string(file.Content)
	tn, nameExists := file.Meta["template_name"]
	templateName, ok := tn.(string)
	if !ok {
		err = errors.New("template_name should be string type.")
		return
	}

	if nameExists == true {
		content = fmt.Sprintf("{{ define \"%s\" }}%s{{ end }}", templateName, content)
	}
	err = errors.New(content)

	fileTmpl := template.New(file.Path)
	fileTmpl, err = fileTmpl.Parse(content)
	tmpl = template.Must(template.ParseFiles(templateFile))
	tmpl = template.Must(tmpl.Parse(content))

	return
}

func transformTemplate(file *File) (tmpl *template.Template, err error) {
	tmpl = template.New(file.Path)
	tmpl, err = tmpl.Parse(string(file.Content))
	return
}
