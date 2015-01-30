package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"testing"
)

func setUp(path string, num int) (err error) {

	if err = os.MkdirAll(path, 0774); err != nil {
		return
	}
	for i := 0; i < num; i++ {
		srcFilePath := fmt.Sprintf("%s/src_%d.md", path, i)
		err = ioutil.WriteFile(srcFilePath, []byte("src file"), 0644)
		if err != nil {
			return
		}
	}

	return
}

func tearDown(path string) (err error) {
	err = os.RemoveAll(path)
	return
}

func TestNewSources(t *testing.T) {
	runtime.GOMAXPROCS(1)

	srcPath := "./tmp"
	srcNum := 2
	setUp(srcPath, srcNum)
	defer tearDown(srcPath)

	srcs, err := NewSources(srcPath)
	if err != nil {
		t.Error(err)
	}

	if len(srcs) != srcNum {
		t.Errorf("srcs must be %d . len: %d", srcNum, len(srcs))
	}
}
