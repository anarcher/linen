package main

const (
	FileTypeMarkdown = iota
	FileTypeTemplate
	FileTypePlain
)

type FileContent map[string]interface{}

type File struct {
	Path    string
	Ext     string
	Type    int
	Content FileContent
}

type Files []*File


func NewFile(path ext string) (*File) {
    file = &File{  Path: path , Ext: ext }
    var fileType int
    switch ext {
    case ".md":
        fileType = FileTypeMarkdown
    case ".tmpl":
        fileType = FileTypeTemplate
    default:
        fileType = FileTypePlain
    }
    return file
}
