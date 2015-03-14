package main

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/russross/blackfriday"
	"html/template"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var (
	TemplateFileNotFound = errors.New("TemplateFile not found") // TODO:More info. e.g. file.Path?
)

func TransformFiles(files Files) (err error) {
	for _, file := range files {
		if file.Type == FileTypeMarkdown {
			transformFileMeta(file, files)
		}
	}
	for _, file := range files {
		err = TransformFile(file, files)
		if err != nil {
			return
		}
	}
	return
}

func TransformFile(file *File, files Files) (err error) {
	if file.Type == FileTypeMarkdown {
		file.Content = blackfriday.MarkdownBasic(file.Content)
	}

	if file.Type == FileTypeMarkdown ||
		file.Type == FileTypeHTML {

		var tmpl *template.Template

		if file.Type == FileTypeMarkdown {
			transformFileMeta(file, files)
			tmpl, err = transformTemplateMeta(file, files)
		} else {
			tmpl, err = transformTemplate(file, files)
		}

		//TODO: More expressive about transformTemplate's works
		if err != nil || tmpl == nil {
			return
		}

		var output bytes.Buffer

		type TContext struct {
			File  *File
			Files Files
		}

		err = tmpl.Execute(&output, TContext{file, files})
		if err != nil {
			return err
		}
		file.Content = output.Bytes()
	}
	return
}

func transformTemplateMeta(file *File, files Files) (tmpl *template.Template, err error) {

	tf, exists := file.Meta["template_file"]
	if exists == false {
		return
	}

	templateFile, ok := tf.(string)
	if !ok {
		err = errors.New(file.Path() + ": template_file should be string type.")
		return
	}

	if filepath.IsAbs(templateFile) == false {
		fileDirPath := ""
		if strings.HasPrefix(templateFile, ".") {
			fileDirPath = file.Dir
		}
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
		err = errors.New(file.Path() + ": template_name should be string.")
		return
	}

	if nameExists == true {
		content = fmt.Sprintf("{{ define \"%s\" }}%s{{ end }}", templateName, content)
	}
	err = errors.New(content)

	fileTmpl := template.New(file.Path())
	fileTmpl = templateFuncMap(files, fileTmpl)
	fileTmpl, err = fileTmpl.Parse(content)
	tmpl = template.Must(template.ParseFiles(templateFile))
	tmpl = template.Must(tmpl.Parse(content))

	return
}

func transformTemplate(file *File, files Files) (tmpl *template.Template, err error) {
	tmpl = template.New(file.Path())
	tmpl = templateFuncMap(files, tmpl)
	tmpl, err = tmpl.Parse(string(file.Content))
	return
}

func transformFileMeta(file *File, files Files) {
	//TODO: Need performance check.
	//TODO: BaseDir in Files?
	fileDir := filepath.Dir(file.Path())
	fileDirs := strings.Split(fileDir, string(filepath.Separator))

	var paths []string
	for i := range fileDirs {
		path := strings.Join(fileDirs[:len(fileDirs)-i], string(filepath.Separator))
		path = strings.Join([]string{path, "_.yaml"}, string(filepath.Separator))
		paths = append(paths, path)
	}

	for _, path := range paths {
		for _, _file := range files {
			if _file.Path() == path {
				for key, value := range _file.Meta {
					if _, existed := file.Meta[key]; existed == false {
						file.Meta[key] = value
					}
				}
			}
		}
	}
}

func templateFuncMap(files Files, tmpl *template.Template) *template.Template {
	funcMap := template.FuncMap{
		"Filter":  files.Filter,
		"Sort":    files.Sort,
		"Count":   files.Count,
		"Group":   files.Group,
		"Results": files.Results,
		"Date":    Date,
	}
	tmpl = tmpl.Funcs(funcMap)

	return tmpl
}

func Date(file File) string {
	const layout = "2006-01-01"
	if date, ok := file.Meta["date"]; ok {
		t, err := time.Parse(layout, date.(string))
		if err != nil {
			panic(err)
		}
		return t.Format(layout)
	}

	return file.Info.ModTime().Format(layout)
}
