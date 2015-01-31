package main

import (
	"runtime"
	"testing"
)

func TestNewSources(t *testing.T) {
	runtime.GOMAXPROCS(1)

	srcPath := "./examples/basic/"
	srcNum := 1

	srcs, err := NewSources(srcPath)
	if err != nil {
		t.Error(err)
	}

	if len(srcs) != 2 {
		t.Errorf("srcs must be %d . len: %d", srcNum, len(srcs))
	}
}
