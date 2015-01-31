package main

import (
	"io/ioutil"
)

func ReadSource(path string) (source Source, err error) {
	var content []byte
	content, err = ioutil.ReadFile(path)
	if err != nil {
		return
	}

	source.path = path
	source.content = string(content)
	return
}
