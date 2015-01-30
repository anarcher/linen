package main

import (
	"flag"
	"log"
	"runtime"
)

var srcPath = flag.String("src", ".", "source path (default: .)")
var destPath = flag.String("dest", "./dest", "dest. path (default: ./dest)")

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	flag.Parse()

	sources, err := NewSources(*srcPath)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(sources)
}
