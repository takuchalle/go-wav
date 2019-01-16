package wav

import (
	"io"
	"os"
)

// Wav is wav file struct.
type Wav struct {
	header WaveHeader
	reader io.ReadSeeker
}

// Open opens Wav file
func Open(name string) (*Wav, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	w := &Wav{}
	w.reader = f

	err = w.Parse()
	if err != nil {
		return nil, err
	}

	return w, err
}
