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

	sources, err := NewSources(*srcPath)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(sources)
}
