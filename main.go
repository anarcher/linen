package main

import (
	"flag"
	"log"
	"runtime"
)

var srcPath = flag.String("src", ".", "source path (default: .)")
var targetPath = flag.String("target", "./target", "target path (default: ./target)")

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

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

	//WriteFiles(files)

}
