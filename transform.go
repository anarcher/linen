package main

import (
	"encoding/json"
	"github.com/russross/blackfriday"
)

func TransformFiles(files Files) (err error) {
	for _, file := range files {

		//Header
		if header, ok := file.Content[FileContentHeader]; ok {
			var v map[string]interface{}
			if err = json.Unmarshal(header.([]byte), &v); err != nil {
				return err
			}
			file.Content[FileContentHeader] = v
		}

		//Body
		if body, ok := file.Content[FileContentBody]; ok {
			html := blackfriday.MarkdownBasic(body.([]byte))
			file.Content[FileContentBody] = html
		}

		if file.Type == FileTypeMarkdown || file.Type == FileTypeTemplate {
			//TODO: Template

		}
	}

	return
}
