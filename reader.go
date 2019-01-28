package wav

import (
	"encoding/binary"
	"io"
)

// Reader is wav reader struct
type Reader struct {
	r io.ReadSeeker

	h waveHeader
}

// NewReader creates new Reader struct.
// Check wav header.
func NewReader(r io.ReadSeeker) (wav *Reader, err error) {
	wav = &Reader{}
	wav.r = r
	err = wav.parseHeader()
	if err != nil {
		return nil, err
	}

	return wav, nil
}

func (wav *Reader) readRiffChunk(buffer []byte) ([]byte, error) {
	if "RIFF" != string(buffer[:4]) {
		return buffer, ErrNoRIFF
	}
	buffer = buffer[4:]

	wav.h.chunkSize = binary.LittleEndian.Uint32(buffer[:4])
	buffer = buffer[4:]

	if "WAVE" != string(buffer[:4]) {
		return buffer, ErrNotWavFile
	}
	buffer = buffer[4:]
	return buffer, nil
}

func (wav *Reader) readFmtSubChunk(buffer []byte) ([]byte, error) {
	if "fmt " != string(buffer[:4]) {
		return buffer, ErrNoFmt
	}
	buffer = buffer[4:]

	wav.h.subChunkSize = binary.LittleEndian.Uint32(buffer[:4])
	buffer = buffer[4:]

	wav.h.audioFormat = binary.LittleEndian.Uint16(buffer[:2])
	buffer = buffer[2:]

	wav.h.numChannels = binary.LittleEndian.Uint16(buffer[:2])
	buffer = buffer[2:]

	wav.h.sampleRate = binary.LittleEndian.Uint32(buffer[:4])
	buffer = buffer[4:]

	wav.h.byteRate = binary.LittleEndian.Uint32(buffer[:4])
	buffer = buffer[4:]

	wav.h.blockAlign = binary.LittleEndian.Uint16(buffer[:2])
	buffer = buffer[2:]

	wav.h.bitsPerSample = binary.LittleEndian.Uint16(buffer[:2])
	buffer = buffer[2:]

	return buffer, nil
}

func (wav *Reader) readDataSubChunk(buffer []byte) []byte {
	buffer = buffer[4:]

	wav.h.subChunk2Size = binary.LittleEndian.Uint32(buffer[:4])
	buffer = buffer[4:]

	return buffer
}

func (wav *Reader) parseHeader() error {
	buffer := make([]byte, headerSize)
	_, err := io.ReadAtLeast(wav.r, buffer, headerSize)
	if err != nil {
		return err
	}

	buffer, err = wav.readRiffChunk(buffer)
	if err != nil {
		return err
	}
	buffer, err = wav.readFmtSubChunk(buffer)
	if err != nil {
		return err
	}
	buffer = wav.readDataSubChunk(buffer)

	/* Reset Read Position */
	wav.r.Seek(0, 0)

	return nil
}

// GetNumChannels returns num of channels
func (wav *Reader) GetNumChannels() uint16 {
	return wav.h.numChannels
}

// GetChunkSize returns chunk size
func (wav *Reader) GetChunkSize() uint32 {
	return wav.h.chunkSize
}

// GetAudioFormat returns audio format
func (wav *Reader) GetAudioFormat() uint16 {
	return wav.h.audioFormat
}

// GetSampleRate returns sample rate
func (wav *Reader) GetSampleRate() uint32 {
	return wav.h.sampleRate
}

// GetByteRate returns byte rate
func (wav *Reader) GetByteRate() uint32 {
	return wav.h.byteRate
}

// GetBlockAlign returns block align
func (wav *Reader) GetBlockAlign() uint16 {
	return wav.h.blockAlign
}

// GetBitsPerSample returns bits per sample
func (wav *Reader) GetBitsPerSample() uint16 {
	return wav.h.bitsPerSample
}

// GetSubChunkSize returns sub chunk size
func (wav *Reader) GetSubChunkSize() uint32 {
	return wav.h.subChunk2Size
}
