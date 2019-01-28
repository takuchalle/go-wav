package wav

import (
	"bytes"
	"io"
	"testing"

	"github.com/cheekybits/is"
)

var (
	riff             = []byte{0x52, 0x49, 0x46, 0x46} // "RIFF"
	chunkSize24      = []byte{0x24, 0x00, 0x00, 0x00} // chunkSize
	wave             = []byte{0x57, 0x41, 0x56, 0x45} // "WAVE"
	fmt20            = []byte{0x66, 0x6d, 0x74, 0x20} // "fmt "
	testRiffChunkFmt = []byte{
		0x10, 0x00, 0x00, 0x00, // LengthOfHeader
		0x01, 0x00, // AudioFormat
		0x01, 0x00, // NumOfChannels
		0x44, 0xac, 0x00, 0x00, // SampleRate
		0x88, 0x58, 0x01, 0x00, // BytesPerSec
		0x02, 0x00, // BytesPerBloc
		0x10, 0x00, // BitsPerSample
		0x64, 0x61, 0x74, 0x61, // "data"
		0x00, 0x00, 0x00, 0x01,
	}
)

func TestParseHeaders_tooShort(t *testing.T) {
	t.Parallel()
	is := is.New(t)

	var b bytes.Buffer
	b.Write(riff)
	wavFile := bytes.NewReader(b.Bytes())
	_, err := NewReader(wavFile)
	is.Err(err)
	is.Equal(err, io.ErrUnexpectedEOF)
}

func TestParseHeaders_brokenRiff(t *testing.T) {
	t.Parallel()
	is := is.New(t)

	var b bytes.Buffer
	b.Write([]byte{0x52, 0x48, 0x46, 0x46}) // broken riff
	b.Write([]byte{0x26, 0x00, 0x00, 0x00}) // chunkSize
	b.Write(wave)
	b.Write(fmt20)
	b.Write(testRiffChunkFmt)
	wavFile := bytes.NewReader(b.Bytes())
	_, err := NewReader(wavFile)
	is.Err(err)
	is.Equal(err, ErrNoRIFF)
}

func TestParseHeaders_brokenFmt(t *testing.T) {
	t.Parallel()
	is := is.New(t)

	var b bytes.Buffer
	b.Write(riff)
	b.Write([]byte{0x26, 0x00, 0x00, 0x00}) // chunkSize
	b.Write(wave)
	b.Write([]byte{0x66, 0x6d, 0x75, 0x20}) // broken fmt
	b.Write(testRiffChunkFmt)
	wavFile := bytes.NewReader(b.Bytes())
	_, err := NewReader(wavFile)
	is.Err(err)
	is.Equal(err, ErrNoFmt)
}

func TestParseHeaders_brokenWave(t *testing.T) {
	t.Parallel()
	is := is.New(t)

	var b bytes.Buffer
	b.Write(riff)
	b.Write([]byte{0x26, 0x00, 0x00, 0x00}) // chunkSize
	b.Write([]byte{0x57, 0x40, 0x56, 0x45}) // broken Wave
	b.Write(fmt20)
	b.Write(testRiffChunkFmt)
	wavFile := bytes.NewReader(b.Bytes())
	_, err := NewReader(wavFile)
	is.Err(err)
	is.Equal(err, ErrNotWavFile)
}
