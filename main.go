package main

import (
	"flag"
	//	"github.com/codegangsta/cli"
	"log"
)

var srcPath = flag.String("src", ".", "source path (default: .)")
var targetPath = flag.String("target", "./target", "target path (default: ./target)")

func main() {

	flag.Parse()

	var err error

	var files Files

	files, err = ReadFiles(*srcPath)
	if err != nil {
		log.Println(err)
	}

	err = TransformFiles(files)
	if err != nil {
		log.Println(err)
	}

	err = WriteFiles(files, *targetPath)
	if err != nil {
		log.Println(err)
	}

}
