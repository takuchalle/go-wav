package waveparser

import (
	"io"
	"log"
	"fmt"
)

type WaveHeader struct {
	chunk_id        int
	chunk_size      int
	format          int
	sub_chunk_id    int
	sub_chunk_size  int
	audio_format    int
	num_channels    int
	sample_rate     int
	byte_rate       int
	block_align     int
	bits_per_sample int
	sub_chunk2_id   int
	sub_chunk2_size int
}

type WaveParser struct {
	header WaveHeader
	reader io.Reader
}

func New(r io.Reader) *WaveParser {
	parser := &WaveParser{}
	parser.reader = r
	return parser
}

func (parser *WaveParser) Parse() {
	buffer := make([]byte , 100)
	_, err := io.ReadAtLeast(parser.reader, buffer, 100)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n",buffer)
}

func (parser *WaveParser) GetHeader() *WaveHeader {
	return &parser.header
}

func (header *WaveHeader) GetChunkId() int {
	return header.chunk_id
}

func (header *WaveHeader) GetChunkSize() int {
	return header.chunk_size
}

