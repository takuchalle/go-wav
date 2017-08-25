package wav

import (
	"io"
)

// Wav is
type Wav struct {
	header WaveHeader
	reader io.ReadSeeker
}

// NewWav creats Wave Parser
func NewWav(r io.ReadSeeker) *Wav {
	parser := &Wav{}
	parser.reader = r
	return parser
}
