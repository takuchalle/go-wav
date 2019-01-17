package main

import (
	"fmt"
	"log"
	"os"

	"github.com/takuyaohashi/go-wav"
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

	w, err := wav.NewReader(f)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Num of channles is %v\n", w.GetNumChannels())
}
