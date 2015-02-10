package main

import (
	"testing"

	"github.com/kr/pretty"
)

func TestReadFiles(t *testing.T) {

	path := "./examples/basic"
	files, err := ReadFiles(path)
	if err != nil {
		t.Error(err)
	}

	if len(files) <= 0 {
		t.Error("files len must be > 0")
	}

	t.Logf("%# v", pretty.Formatter(files))

}
