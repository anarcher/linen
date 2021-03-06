package main

import (
	//	"github.com/kr/pretty"
	"os"
	"testing"
)

func TestWriteFiles(t *testing.T) {
	var err error

	srcPath := "./examples/template/"
	destPath := "/tmp/linen-test01/"

	if err = os.RemoveAll(destPath); err != nil {
		t.Error(err)
	}

	files, err := ReadFiles(srcPath)
	if err != nil {
		t.Error(err)
	}

	err = TransformFiles(files)
	if err != nil {
		t.Error(err)
	}

	err = WriteFiles(files, destPath)
	if err != nil {
		t.Error(err)
	}

}
