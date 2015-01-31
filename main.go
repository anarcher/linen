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

	sources := Sources{}

	err = Read(sources, *srcPath)
	if err != nil {
		log.Println(err)
		return
	}

	err = Transform(sources)
	if err != nil {
		log.Println(err)
		return
	}
	//Write(sources)

}
