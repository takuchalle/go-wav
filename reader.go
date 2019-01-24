package wav

import (
	"encoding/binary"
	"io"
)

type Reader struct {
	r io.ReadSeeker

	h WaveHeader
}

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

func (wav *Reader) GetNumChannels() uint16 {
	return wav.h.NumChannels
}

func (wav *Reader) GetChunkSize() uint32 {
	return wav.h.ChunkSize
}

func (wav *Reader) GetSubChunkSize() uint32 {
	return wav.h.SubChunkSize
}

func (wav *Reader) GetAudioFormat() uint16 {
	return wav.h.AudioFormat
}

func (wav *Reader) GetSampleRate() uint32 {
	return wav.h.SampleRate
}

func (wav *Reader) GetByteRate() uint32 {
	return wav.h.ByteRate
}

func (wav *Reader) GetBlockAlign() uint32 {
	return wav.h.BlockAlign
}

func (wav *Reader) GetBitsPerSample() uint32 {
	return wav.h.BitsPerSample
}

func (wav *Reader) GetSubChunkSize() uint32 {
	return wav.h.SubChunk2Size
}

