package wav

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

type WriterParam struct {
	SampleRate    uint32
	BitsPerSample uint16
	NumChannels   uint16
	AudioFormat   AudioFormat
}

type Writer struct {
	r io.WriteSeeker

	p WriterParam
}

func NewWriter(r io.WriteSeeker, p WriterParam) (w *Writer, err error) {
	if p.SampleRate == 0 || p.BitsPerSample == 0 || p.NumChannels == 0 {
		return nil, ErrInvalidFmt
	}

	w = &Writer{r: r, p: p}
	w.writeHeader()
	return w, nil
}

func (w *Writer) Close() {

}

func (w *Writer) writeRiffChunk(buf *bytes.Buffer) error {
	header := fmt.Sprintf("RIFF")
	err := binary.Write(buf, binary.LittleEndian, []byte(header))
	if err != nil {
		return ErrFailedWrite
	}

	// Chunk Size
	err = binary.Write(buf, binary.LittleEndian, float32(0))
	if err != nil {
		return ErrFailedWrite
	}

	header = fmt.Sprintf("WAVE")
	err = binary.Write(buf, binary.LittleEndian, []byte(header))
	if err != nil {
		return ErrFailedWrite
	}
	return nil
}

func (w *Writer) writeFmtSubChunk(buf *bytes.Buffer) error {

	header := fmt.Sprintf("fmt ")
	err := binary.Write(buf, binary.LittleEndian, []byte(header))
	if err != nil {
		return ErrFailedWrite
	}

	// SubChunk1Size
	err = binary.Write(buf, binary.LittleEndian, int32(00000))
	if err != nil {
		return ErrFailedWrite
	}

	// AudioFormat
	err = binary.Write(buf, binary.LittleEndian, int16(1))
	if err != nil {
		return ErrFailedWrite
	}

	// Num Channels
	err = binary.Write(buf, binary.LittleEndian, int16(w.p.NumChannels))
	if err != nil {
		return ErrFailedWrite
	}

	// Sample Rate
	err = binary.Write(buf, binary.LittleEndian, int32(w.p.SampleRate))
	if err != nil {
		return ErrFailedWrite
	}

	// Byte Rate
	byteRate := w.p.SampleRate * uint32(w.p.NumChannels) * uint32(w.p.BitsPerSample) / 8
	err = binary.Write(buf, binary.LittleEndian, int32(byteRate))
	if err != nil {
		return ErrFailedWrite
	}

	// Block Align
	blockAlign := uint32(w.p.NumChannels) * uint32(w.p.BitsPerSample) / 8
	err = binary.Write(buf, binary.LittleEndian, int16(blockAlign))
	if err != nil {
		return ErrFailedWrite
	}

	// Bits Per Sample
	err = binary.Write(buf, binary.LittleEndian, int16(w.p.BitsPerSample))
	if err != nil {
		return ErrFailedWrite
	}

	// data
	header = fmt.Sprintf("data")
	err = binary.Write(buf, binary.LittleEndian, []byte(header))
	if err != nil {
		return ErrFailedWrite
	}

	// SubChunk2Size
	err = binary.Write(buf, binary.LittleEndian, int32(4))
	if err != nil {
		return ErrFailedWrite
	}

	return nil
}

func (w *Writer) writeHeader() {
	buf := new(bytes.Buffer)

	err := w.writeRiffChunk(buf)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}

	err = w.writeFmtSubChunk(buf)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}

	w.r.Write(buf.Bytes())
}

func (w *Writer) WriteSamples() {

}
