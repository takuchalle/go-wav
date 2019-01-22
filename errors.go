package wav

import (
	"errors"
)

var (
	ErrNoRIFF     = errors.New("No RIFF word")
	ErrNotWavFile = errors.New("Not a Wav file")
	ErrNoFmt      = errors.New("No fmt word")
)
