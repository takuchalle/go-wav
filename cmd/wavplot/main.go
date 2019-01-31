package main

import (
	"log"
	"os"

	"github.com/takuyaohashi/go-wav"
	"github.com/takuyaohashi/go-wav/wavplotter"
)

func usage() {
	log.Fatal("Usage: wavplot [wave file name]")
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

	plotter := wavplotter.NewPlotter(w)
	plotter.Plot()
}
