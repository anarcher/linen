package main

import (
	//	"github.com/kr/pretty"
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

	//pretty.Print(srcs)

}
