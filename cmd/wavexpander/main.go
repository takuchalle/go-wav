package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/takuyaohashi/go-wav"
)

func usage() {
	fmt.Println("Usage: go-wavexpand [wav file]")
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

	parser := wav.NewWav(f)
	err = parser.Parse()
	if err != nil {
		log.Fatal(err)
	}
	header := parser.GetHeader()

	fmt.Printf("size = %d\n", header.SubChunk2Size)
	buffer := make([]byte, header.SubChunk2Size)
	_, err = io.ReadAtLeast(f, buffer, int(header.SubChunk2Size))
	wf, err2 := os.Create("hoge.wav")
	if err2 != nil {
		log.Fatal(err)
	}
	defer wf.Close()
	writer := bufio.NewWriter(wf)
	writer.Write(buffer)
}
