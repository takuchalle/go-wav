package wav

const (
	headerSize = 44
)

// AudioFormat describes audio format
type AudioFormat int

const (
	// AudioFormatPCM describes wav file format is pcm.
	AudioFormatPCM AudioFormat = iota

	// AudioFormatBitstream describes wav file format is bitstream.
	AudioFormatBitstream
)

// String shows audio format
func (af AudioFormat) String() string {
	switch af {
	case AudioFormatPCM:
		return "PCM"
	case AudioFormatBitstream:
		return "BitStream"
	default:
		return "Error"
	}
}

// WaveHeader is wave header struct
type waveHeader struct {
	chunkSize     uint32
	subChunkSize  uint32
	audioFormat   uint16
	numChannels   uint16
	sampleRate    uint32
	byteRate      uint32
	blockAlign    uint16
	bitsPerSample uint16
	subChunk2Size uint32
}
