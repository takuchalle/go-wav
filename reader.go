package wav

import (
	"encoding/binary"
	"io"
)

// Reader is wav reader struct
type Reader struct {
	r io.ReadSeeker

	h WaveHeader
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

	wav.h.ChunkSize = binary.LittleEndian.Uint32(buffer[:4])
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

	wav.h.SubChunkSize = binary.LittleEndian.Uint32(buffer[:4])
	buffer = buffer[4:]

	wav.h.AudioFormat = binary.LittleEndian.Uint16(buffer[:2])
	buffer = buffer[2:]

	wav.h.NumChannels = binary.LittleEndian.Uint16(buffer[:2])
	buffer = buffer[2:]

	wav.h.SampleRate = binary.LittleEndian.Uint32(buffer[:4])
	buffer = buffer[4:]

	wav.h.ByteRate = binary.LittleEndian.Uint32(buffer[:4])
	buffer = buffer[4:]

	wav.h.BlockAlign = binary.LittleEndian.Uint16(buffer[:2])
	buffer = buffer[2:]

	wav.h.BitsPerSample = binary.LittleEndian.Uint16(buffer[:2])
	buffer = buffer[2:]

	return buffer, nil
}

func (wav *Reader) readDataSubChunk(buffer []byte) []byte {
	buffer = buffer[4:]

	wav.h.SubChunk2Size = binary.LittleEndian.Uint32(buffer[:4])
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
	return wav.h.NumChannels
}

// GetChunkSize returns chunk size
func (wav *Reader) GetChunkSize() uint32 {
	return wav.h.ChunkSize
}

// GetAudioFormat returns audio format
func (wav *Reader) GetAudioFormat() uint16 {
	return wav.h.AudioFormat
}

// GetSampleRate returns sample rate
func (wav *Reader) GetSampleRate() uint32 {
	return wav.h.SampleRate
}

// GetByteRate returns byte rate
func (wav *Reader) GetByteRate() uint32 {
	return wav.h.ByteRate
}

// GetBlockAlign returns block align
func (wav *Reader) GetBlockAlign() uint16 {
	return wav.h.BlockAlign
}

// GetBitsPerSample returns bits per sample
func (wav *Reader) GetBitsPerSample() uint16 {
	return wav.h.BitsPerSample
}

// GetSubChunkSize returns sub chunk size
func (wav *Reader) GetSubChunkSize() uint32 {
	return wav.h.SubChunk2Size
}
