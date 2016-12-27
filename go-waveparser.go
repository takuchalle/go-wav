package waveparser

import (
	"io"
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
}

func New(r io.Reader) *WaveParser {
	header := &WaveParser{}
	return header
}

func (parser *WaveParser) Parse() {
	
}

func (parser *WaveParser) GetHeader() *WaveHeader {
	return &parser.header
}

func (header *WaveHeader) GetChunkId() int {
	return header.chunk_id
}
