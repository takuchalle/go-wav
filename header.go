package wav

const (
	headerSize = 44
)

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
