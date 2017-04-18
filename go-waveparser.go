package waveparser

import (
	"io"
	"log"
	"encoding/binary"
)

const HEADER_SIZE = 44

type WaveHeader struct {
	chunk_id        string
	chunk_size      uint32
	format          string
	sub_chunk_id    string
	sub_chunk_size  uint32
	audio_format    uint16
	num_channels    uint16
	sample_rate     uint32
	byte_rate       uint32
	block_align     uint16
	bits_per_sample uint16
	sub_chunk2_id   string
	sub_chunk2_size uint32
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

func (parser *WaveParser) ReadRiffChunk(buffer []byte) []byte {
	if "RIFF" != string(buffer[:4]) {
		log.Fatal("This is not WAVE file!\n")
	}
	parser.header.chunk_id = string(buffer[:4])
	buffer = buffer[4:]

	parser.header.chunk_size = binary.LittleEndian.Uint32(buffer[:4])
	buffer = buffer[4:]

	if "WAVE" != string(buffer[:4]) {
		log.Fatal("This is not WAVE file!\n")
	}
	parser.header.format = string(buffer[:4])
	buffer = buffer[4:]
	return buffer
}

func (parser *WaveParser) ReadFmtSubChunk(buffer []byte) []byte {
	if "fmt " != string(buffer[:4]) {
		log.Fatal("This is not WAVE file!\n")
	}
	parser.header.sub_chunk_id = string(buffer[:4])
	buffer = buffer[4:]

	parser.header.sub_chunk_size = binary.LittleEndian.Uint32(buffer[:4])
	buffer = buffer[4:]

	parser.header.audio_format = binary.LittleEndian.Uint16(buffer[:2])
	buffer = buffer[2:]

	parser.header.num_channels = binary.LittleEndian.Uint16(buffer[:2])
	buffer = buffer[2:]

	parser.header.sample_rate = binary.LittleEndian.Uint32(buffer[:4])
	buffer = buffer[4:]

	parser.header.byte_rate = binary.LittleEndian.Uint32(buffer[:4])
	buffer = buffer[4:]

	parser.header.block_align = binary.LittleEndian.Uint16(buffer[:2])
	buffer = buffer[2:]

	parser.header.bits_per_sample = binary.LittleEndian.Uint16(buffer[:2])
	buffer = buffer[2:]
	
	return buffer
}

func (parser *WaveParser) ReadDataSubChunk(buffer []byte) []byte {

	parser.header.sub_chunk2_id = string(buffer[:4])
	buffer = buffer[4:]

	parser.header.sub_chunk2_size = binary.LittleEndian.Uint32(buffer[:4])
	buffer = buffer[4:]
	
	return buffer
}

func (parser *WaveParser) Parse() {
	buffer := make([]byte, HEADER_SIZE)
	_, err := io.ReadAtLeast(parser.reader, buffer, HEADER_SIZE)
	if err != nil {
		log.Fatal(err)
	}

	buffer = parser.ReadRiffChunk(buffer)
	buffer = parser.ReadFmtSubChunk(buffer)
	buffer = parser.ReadDataSubChunk(buffer)
}

func (parser *WaveParser) GetHeader() *WaveHeader {
	return &parser.header
}

func (header *WaveHeader) GetChunkId() string {
	return header.chunk_id
}

func (header *WaveHeader) GetChunkSize() uint32 {
	return header.chunk_size
}

func (header *WaveHeader) GetFormat() string {
	return header.format
}

func (header *WaveHeader) GetSubChunkId() string {
	return header.sub_chunk_id
}

func (header *WaveHeader) GetSubChunkSize() uint32 {
	return header.sub_chunk_size
}

func (header *WaveHeader) GetAudioFormat() uint16 {
	return header.audio_format
}

func (header *WaveHeader) GetNumChannels() uint16 {
	return header.num_channels
}

func (header *WaveHeader) GetSampleRate() uint32 {
	return header.sample_rate
}

func (header *WaveHeader) GetByteRate() uint32 {
	return header.byte_rate
}

func (header *WaveHeader) GetBlockAlign() uint16 {
	return header.block_align
}

func (header *WaveHeader) GetBitPerSample() uint16 {
	return header.bits_per_sample
}

func (header *WaveHeader) GetSubChunk2Id() string {
	return header.sub_chunk2_id
}

func (header *WaveHeader) GetSubChunk2Size() uint32 {
	return header.sub_chunk2_size
}

