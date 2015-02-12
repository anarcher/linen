package main

import (
	"bytes"
	"encoding/json"
	"github.com/russross/blackfriday"
	"html/template"
)

func TransformFiles(files Files) (err error) {
	for _, file := range files {

		//Header
		if header, ok := file.Meta[FileHeaderRaw]; ok {
			var v map[string]interface{}
			if err = json.Unmarshal(header.([]byte), &v); err != nil {
				return err
			}
			file.Meta = v
		}

		//Body
		if file.Type == FileTypeMarkdown {
			file.Content = blackfriday.MarkdownBasic(file.Content)
		}

		if file.Type == FileTypeMarkdown || file.Type == FileTypeTemplate {

			tmpl := template.New(file.Path)
			tmpl, err = tmpl.Parse(string(file.Content))
			if err != nil {
				return err
			}

			var output bytes.Buffer

			err = tmpl.Execute(&output, file.Meta)
			if err != nil {
				return err
			}
			file.Content = output.Bytes()
		}
	}

	return
}
