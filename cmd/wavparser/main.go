package main

import (
	"fmt"
	"log"
	"os"

	"github.com/takuchalle/go-wav"
)

func usage() {
	log.Fatal("Usage: waveparser [wave file name]")
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

	w := wav.NewReader(f)

	w.Parse()

	fmt.Printf("Chunk Size:\t%v\n", w.GetChunkSize())
	fmt.Printf("SubChunk Size:\t%v\n", w.GetSubChunkSize())
	fmt.Printf("Audio Format:\t%v\n", w.GetAudioFormat())
	fmt.Printf("Channles:\t%v\n", w.GetNumChannels())
	fmt.Printf("SampleRate:\t%v\n", w.GetSampleRate())
	fmt.Printf("ByteRate:\t%v\n", w.GetByteRate())
	fmt.Printf("BitsPerSample:\t%v\n", w.GetBitsPerSample())
	fmt.Printf("BlockAlign:\t%v\n", w.GetBlockAlign())
}
