package wav

const (
	headerSize = 44
)

// WaveHeader is wave header struct
type WaveHeader struct {
	ChunkSize     uint32
	SubChunkSize  uint32
	AudioFormat   uint16
	NumChannels   uint16
	SampleRate    uint32
	ByteRate      uint32
	BlockAlign    uint16
	BitsPerSample uint16
	SubChunk2Size uint32
}
