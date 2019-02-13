package wav

import (
	"errors"
)

var (
	ErrNoRIFF      = errors.New("No RIFF word")
	ErrNotWavFile  = errors.New("Not a Wav file")
	ErrNoFmt       = errors.New("No fmt word")
	ErrInvalidFmt  = errors.New("Audio Format Error")
	ErrFailedWrite = errors.New("Failet to write")
)
