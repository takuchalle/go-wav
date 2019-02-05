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
	r io.Writer

	p WriterParam
}

func NewWriter(r io.Writer, p WriterParam) (w *Writer, err error) {
	w = &Writer{r: r, p: p}
	return w, nil
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
