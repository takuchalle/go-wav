package wav

import (
	"errors"
)

var (
	// ErrNoRIFF is the error about missing RIFF word.
	ErrNoRIFF = errors.New("No RIFF word")

	// ErrNotWavFile is the error that input file is not wav file.
	ErrNotWavFile = errors.New("Not a Wav file")

	// ErrNoFmt is the error about missing WAV word.
	ErrNoFmt = errors.New("No fmt word")

	// ErrInvalidFmt is the error about fomat error.
	ErrInvalidFmt = errors.New("Audio Format Error")

	// ErrFailedWrite is the error about failing to write.
	ErrFailedWrite = errors.New("Failet to write")
)
