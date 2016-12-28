package main

import (
	"fmt"
	"github.com/takuyaohashi/go-waveparser"
	"log"
	"os"
)

func usage() {
	log.Fatal("Usage: go-waveparser [wave file name]")
}

func main() {
	if len(os.Args) != 2 {
		usage()
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	parser := waveparser.New(f)
	parser.Parse()
	header := parser.GetHeader()
	fmt.Printf("chunk_id = %d\n", header.GetChunkId())
}
