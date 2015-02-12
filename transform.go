package main

import (
	"encoding/json"
	"github.com/russross/blackfriday"
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
			//TODO: Template

		}
	}

	return
}
