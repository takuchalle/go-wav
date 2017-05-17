package waveparser

import (
	"encoding/binary"
	"io"
	"log"
)

// HeaderSize is Wav Header Size
const HeaderSize = 44

// WaveHeader is wave header
type WaveHeader struct {
	ChunkID       string
	ChunkSize     uint32
	Format        string
	SubChunkID    string
	SubChunkSize  uint32
	AudioFormat   uint16
	NumChannels   uint16
	SampleRate    uint32
	ByteRate      uint32
	BlockAlign    uint16
	BitsPerSample uint16
	SubChunk2ID   string
	SubChunk2Size uint32
}

// WaveParser is 
type WaveParser struct {
	header WaveHeader
	reader io.Reader
}

// New creats Wave Parser
func New(r io.Reader) *WaveParser {
	parser := &WaveParser{}
	parser.reader = r
	return parser
}

func (parser *WaveParser) readRiffChunk(buffer []byte) []byte {
	if "RIFF" != string(buffer[:4]) {
		log.Fatal("This is not WAVE file!\n")
	}
	parser.header.ChunkID = string(buffer[:4])
	buffer = buffer[4:]

	parser.header.ChunkSize = binary.LittleEndian.Uint32(buffer[:4])
	buffer = buffer[4:]

	if "WAVE" != string(buffer[:4]) {
		log.Fatal("This is not WAVE file!\n")
	}
	parser.header.Format = string(buffer[:4])
	buffer = buffer[4:]
	return buffer
}

func (parser *WaveParser) readFmtSubChunk(buffer []byte) []byte {
	if "fmt " != string(buffer[:4]) {
		log.Fatal("This is not WAVE file!\n")
	}
	parser.header.SubChunkID = string(buffer[:4])
	buffer = buffer[4:]

	parser.header.SubChunkSize = binary.LittleEndian.Uint32(buffer[:4])
	buffer = buffer[4:]

	parser.header.AudioFormat = binary.LittleEndian.Uint16(buffer[:2])
	buffer = buffer[2:]

	parser.header.NumChannels = binary.LittleEndian.Uint16(buffer[:2])
	buffer = buffer[2:]

	parser.header.SampleRate = binary.LittleEndian.Uint32(buffer[:4])
	buffer = buffer[4:]

	parser.header.ByteRate = binary.LittleEndian.Uint32(buffer[:4])
	buffer = buffer[4:]

	parser.header.BlockAlign = binary.LittleEndian.Uint16(buffer[:2])
	buffer = buffer[2:]

	parser.header.BitsPerSample = binary.LittleEndian.Uint16(buffer[:2])
	buffer = buffer[2:]

	return buffer
}

func (parser *WaveParser) readDataSubChunk(buffer []byte) []byte {

	parser.header.SubChunk2ID = string(buffer[:4])
	buffer = buffer[4:]

	parser.header.SubChunk2Size = binary.LittleEndian.Uint32(buffer[:4])
	buffer = buffer[4:]

	return buffer
}

func (parser *WaveParser) Parse() {
	buffer := make([]byte, HeaderSize)
	_, err := io.ReadAtLeast(parser.reader, buffer, HeaderSize)
	if err != nil {
		log.Fatal(err)
	}

	buffer = parser.readRiffChunk(buffer)
	buffer = parser.readFmtSubChunk(buffer)
	buffer = parser.readDataSubChunk(buffer)
}

func (parser *WaveParser) GetHeader() *WaveHeader {
	return &parser.header
}
