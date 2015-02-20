package main

import (
	"github.com/kr/pretty"
	"path/filepath"
	"testing"
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

func TestReadYAMLConf(t *testing.T) {
	path := "./examples/basic"

	files, err := ReadFiles(path)
	if err != nil {
		t.Error(err)
	}

	if len(files) <= 0 {
		t.Error("files len must be > 0")
	}

	for _, f := range files {
		if filepath.Base(f.Path) == "_.yaml" {
			if len(f.Meta) <= 0 {
				t.Errorf("%s could not read", f.Path)
			}
		}
	}

}
