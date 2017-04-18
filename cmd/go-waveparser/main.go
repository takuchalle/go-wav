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
	fmt.Printf("Format \t\t= %s\n", header.GetFormat())
	fmt.Printf("Sub Chunk Size \t= %d\n", header.GetSubChunkSize())
	fmt.Printf("Audio Format \t= %d\n", header.GetAudioFormat())
	fmt.Printf("Num Of Channels = %d\n", header.GetNumChannels())
	fmt.Printf("Sample Rate \t= %d\n", header.GetSampleRate())
	fmt.Printf("Byte Rate \t= %d\n", header.GetByteRate())
	fmt.Printf("Block Align \t= %d\n", header.GetBlockAlign())
	fmt.Printf("Bit Per Sample \t= %d\n", header.GetBitPerSample())
	fmt.Printf("Data Size \t= %d\n", header.GetSubChunk2Size())
}

