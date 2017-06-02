package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/takuyaohashi/go-wav"
)

func usage() {
	fmt.Println("Usage: go-wavexpand [wav file]")
}

func expand(in <-chan []byte) <-chan []byte {
	out := make(chan []byte)

	go func() {
		defer close(out)
		for data := range in {
			buf := make([]byte, 4)
			for i := 0; i < len(data); i++ {
				buf[i+1] = data[i]
			}
			out <- buf
		}
	}()

	return out
}

func read(reader *bufio.Reader) <-chan []byte {
	out := make(chan []byte)

	go func() {
		defer close(out)
		for {
			buf := make([]byte, 3)
			_, err := reader.Read(buf)
			if err == io.EOF {
				break
			}
			out <- buf
		}
	}()
	return out
}

func run() int {
	if len(os.Args) != 2 {
		usage()
		return 1
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to open file\n`")
		return 1
	}
	defer f.Close()

	parser := wav.NewWav(f)
	err = parser.Parse()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to parse")
		return 1
	}
	header := parser.GetHeader()

	wf, err2 := os.Create("hoge.wav")
	if err2 != nil {
		fmt.Fprintln(os.Stderr, "Failed to create new file")
		return 1
	}
	defer wf.Close()
	writer := bufio.NewWriter(wf)

	reader := bufio.NewReaderSize(f, int(header.SubChunk2Size))
	out := read(reader)
	expand := expand(out)

	for i := range expand {
		writer.Write(i)
	}

	writer.Flush()

	return 0
}

func main() {
	if run() != 0 {
		os.Exit(1)
	}
}
