package main

import (
	"github.com/kr/pretty"
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
