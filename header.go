package wav

import (
	"encoding/binary"
	"errors"
	"io"
	"log"
)

const headerSize = 44

// WaveHeader is wave header
type WaveHeader struct {
	ChunkSize     uint32
	SubChunkSize  uint32
	AudioFormat   uint16
	NumChannels   uint16
	SampleRate    uint32
	ByteRate      uint32
	BlockAlign    uint16
	BitsPerSample uint16
	SubChunk2Size uint32
}

func (parser *Wav) readRiffChunk(buffer []byte) ([]byte, error) {
	if "RIFF" != string(buffer[:4]) {
		return buffer, errors.New("This is not wav file")
	}
	buffer = buffer[4:]

	parser.header.ChunkSize = binary.LittleEndian.Uint32(buffer[:4])
	buffer = buffer[4:]

	if "WAVE" != string(buffer[:4]) {
		log.Fatal("This is not WAVE file!\n")
	}
	buffer = buffer[4:]
	return buffer, nil
}

func (parser *Wav) readFmtSubChunk(buffer []byte) ([]byte, error) {
	if "fmt " != string(buffer[:4]) {
		return buffer, errors.New("This is not wav file")
	}
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
	buffer = buffer[4:]

	parser.header.SubChunk2Size = binary.LittleEndian.Uint32(buffer[:4])
	buffer = buffer[4:]

	return buffer
}

func (parser *Wav) Parse() error {
	buffer := make([]byte, headerSize)
	_, err := io.ReadAtLeast(parser.reader, buffer, headerSize)
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

	/* Reset Read Position */
	parser.reader.Seek(0, 0)

	return nil
}

func (parser *Wav) GetHeader() *WaveHeader {
	return &parser.header
}
