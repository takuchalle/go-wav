package waveparser

type WaveHeader struct {
	chunk_id        int
	chunk_size      int
	format          int
	sub_chunk_id    int
	sub_chunk_size  int
	audio_format    int
	num_channels    int
	sample_rate     int
	byte_rate       int
	block_align     int
	bits_per_sample int
	sub_chunk2_id   int
	sub_chunk2_size int
}

func New() *WaveHeader {
	header := &WaveHeader{}
	return header
}
