package wav

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

type WriterParam struct {
	sampleRate    uint32
	bitsPerSample uint16
	numChannels   uint16
	audioFormat   AudioFormat
}

type Writer struct {
	r io.Writer
}

func NewWriter(r io.Writer, p WriterParam) (w *Writer, err error) {
	w = &Writer{}
	w.r = r
	w.setDefaultFormat()
	return w, nil
}

func (w *Writer) setDefaultFormat() {
	w.SetFormat(WriterParam{sampleRate: 1, bitsPerSample: 1, numChannels: 1, audioFormat: AudioFormatPCM})
}

func (w *Writer) SetFormat(param WriterParam) {

}

func (w *Writer) writeHeader() {
	buf := new(bytes.Buffer)
	header := fmt.Sprintf("RIFFWAVEfmt")
	err := binary.Write(buf, binary.LittleEndian, []byte(header))
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	fmt.Printf("%x\n", buf.Bytes())
	w.r.Write(buf.Bytes())
}

func (w *Writer) Write() {
	w.writeHeader()
}
