package main

type Path string

type Source map[string]interface{}

type Sources map[Path]Source

const (
	SourceTypeMarkdown = iota
	SourceTypeTemplate
	SourceTypeCopy
)
