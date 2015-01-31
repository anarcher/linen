package main

import (
	"github.com/kr/pretty"
	"testing"
)

func TestReadSourcesAll(t *testing.T) {
	var err error

	srcs := Sources{}
	path := "./examples/basic"
	err = Read(srcs, path)

	if err != nil {
		t.Error(err)
	}

	if len(srcs) != 2 {
		t.Error("srcs size is not corrent")
	}

	t.Logf("%# v", pretty.Formatter(srcs))
}

func TestreadHeaderAndBody(t *testing.T) {
	var err error

	path := "./examples/basic/post1.md"
	src := Source{}
	err = ReadOne(src, path)

	if err != nil {
		t.Error(err)
	}

}
