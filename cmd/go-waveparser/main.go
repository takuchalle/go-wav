package main

import (
	"fmt"
	"os"
	"log"
	"github.com/takuyaohashi/go-waveparser"
)

func usage() {
	fmt.Printf("Usage: go-waveparser [wave file name]\n");
	os.Exit(1)
}

func main() {
	if(len(os.Args) != 2) { usage() }
	
	f, err := os.Open(os.Args[1]) // For read access.
	if err != nil {
		log.Fatal(err)
	}
	
	parser := waveparser.New(f)
	parser.Parse()
	header := parser.GetHeader()
	fmt.Printf("chunk_id = %d\n", header.GetChunkId());
}
