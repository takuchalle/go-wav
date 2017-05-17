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
	chunkID       string
	chunkSize     uint32
	format        string
	subChunkID    string
	subChunkSize  uint32
	audioFormat   uint16
	numChannels   uint16
	sampleRate    uint32
	byteRate      uint32
	blockAlign    uint16
	bitsPerSample uint16
	subChunk2ID   string
	subChunk2Size uint32
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
	parser.header.chunkID = string(buffer[:4])
	buffer = buffer[4:]

	parser.header.chunkSize = binary.LittleEndian.Uint32(buffer[:4])
	buffer = buffer[4:]

	if "WAVE" != string(buffer[:4]) {
		log.Fatal("This is not WAVE file!\n")
	}
	parser.header.format = string(buffer[:4])
	buffer = buffer[4:]
	return buffer
}

func (parser *WaveParser) readFmtSubChunk(buffer []byte) []byte {
	if "fmt " != string(buffer[:4]) {
		log.Fatal("This is not WAVE file!\n")
	}
	parser.header.subChunkID = string(buffer[:4])
	buffer = buffer[4:]

	parser.header.subChunkSize = binary.LittleEndian.Uint32(buffer[:4])
	buffer = buffer[4:]

	parser.header.audioFormat = binary.LittleEndian.Uint16(buffer[:2])
	buffer = buffer[2:]

	parser.header.numChannels = binary.LittleEndian.Uint16(buffer[:2])
	buffer = buffer[2:]

	parser.header.sampleRate = binary.LittleEndian.Uint32(buffer[:4])
	buffer = buffer[4:]

	parser.header.byteRate = binary.LittleEndian.Uint32(buffer[:4])
	buffer = buffer[4:]

	parser.header.blockAlign = binary.LittleEndian.Uint16(buffer[:2])
	buffer = buffer[2:]

	parser.header.bitsPerSample = binary.LittleEndian.Uint16(buffer[:2])
	buffer = buffer[2:]

	return buffer
}

func (parser *WaveParser) readDataSubChunk(buffer []byte) []byte {

	parser.header.subChunk2ID = string(buffer[:4])
	buffer = buffer[4:]

	parser.header.subChunk2Size = binary.LittleEndian.Uint32(buffer[:4])
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

func (header *WaveHeader) GetChunkID() string {
	return header.chunkID
}

func (header *WaveHeader) GetChunkSize() uint32 {
	return header.chunkSize
}

func (header *WaveHeader) GetFormat() string {
	return header.format
}

func (header *WaveHeader) GetSubChunkID() string {
	return header.subChunkID
}

func (header *WaveHeader) GetSubChunkSize() uint32 {
	return header.subChunkSize
}

func (header *WaveHeader) GetAudioFormat() uint16 {
	return header.audioFormat
}

func (header *WaveHeader) GetNumChannels() uint16 {
	return header.numChannels
}

func (header *WaveHeader) GetSampleRate() uint32 {
	return header.sampleRate
}

func (header *WaveHeader) GetByteRate() uint32 {
	return header.byteRate
}

func (header *WaveHeader) GetBlockAlign() uint16 {
	return header.blockAlign
}

func (header *WaveHeader) GetBitPerSample() uint16 {
	return header.bitsPerSample
}

func (header *WaveHeader) GetSubChunk2Id() string {
	return header.subChunk2ID
}

func (header *WaveHeader) GetSubChunk2Size() uint32 {
	return header.subChunk2Size
}
