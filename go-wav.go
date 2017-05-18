package wav

import (
	"encoding/binary"
	"errors"
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

// Wav is
type Wav struct {
	header WaveHeader
	reader io.Reader
}

// New creats Wave Parser
func New(r io.Reader) *Wav {
	parser := &Wav{}
	parser.reader = r
	return parser
}

func (parser *Wav) readRiffChunk(buffer []byte) ([]byte, error) {
	if "RIFF" != string(buffer[:4]) {
		return buffer, errors.New("This is not wav file")
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
	return buffer, nil
}

func (parser *Wav) readFmtSubChunk(buffer []byte) ([]byte, error) {
	if "fmt " != string(buffer[:4]) {
		return buffer, errors.New("This is not wav file")
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

	return buffer, nil
}

func (parser *Wav) readDataSubChunk(buffer []byte) []byte {

	parser.header.SubChunk2ID = string(buffer[:4])
	buffer = buffer[4:]

	parser.header.SubChunk2Size = binary.LittleEndian.Uint32(buffer[:4])
	buffer = buffer[4:]

	return buffer
}

func (parser *Wav) Parse() error {
	buffer := make([]byte, HeaderSize)
	_, err := io.ReadAtLeast(parser.reader, buffer, HeaderSize)
	if err != nil {
		return err
	}

	buffer, err = parser.readRiffChunk(buffer)
	if err != nil {
		return err
	}
	buffer, err = parser.readFmtSubChunk(buffer)
	if err != nil {
		return err
	}
	buffer = parser.readDataSubChunk(buffer)

	return nil
}

func (parser *Wav) GetHeader() *WaveHeader {
	return &parser.header
}
