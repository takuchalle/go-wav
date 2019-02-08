package main

import (
	"math"
	"os"

	"github.com/takuyaohashi/go-wav"
	"github.com/urfave/cli"
)

func generate_sin(c *cli.Context, w *wav.Writer) error {

	length := c.Float64("sec") * c.Float64("rate")

	for i := 0; i < int(length); i += 256 {
		samples := make([]int16, 256)
		for j := 0; j < 256; j++ {
			samples[j] = int16(math.Sin(math.Pi*2.0*float64(i+j)/(c.Float64("f"))) * float64(math.MaxInt16))
		}
		w.WriteSamples(samples)
	}

	return nil
}

func generate(c *cli.Context) error {
	f, err := os.Create(c.String("output"))
	if err != nil {
		return cli.NewExitError(err, 1)
	}
	defer f.Close()

	samplerate := uint32(c.Int("rate"))
	w, err := wav.NewWriter(f, wav.WriterParam{SampleRate: samplerate, BitsPerSample: 16, NumChannels: 1, AudioFormat: wav.AudioFormatPCM})
	if err != nil {
		return cli.NewExitError(err, 1)
	}

	generate_sin(c, w)

	w.Close()

	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "Sin Wav Genernator"
	app.Usage = "Generate Sin Wav"
	app.Version = "0.0.1"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "output, o",
			Value: "sin.wav",
			Usage: "output file name",
		},

		cli.Float64Flag{
			Name:  "sec, s",
			Value: 5.0,
			Usage: "output sample length",
		},

		cli.IntFlag{
			Name:  "freq, f",
			Value: 100,
			Usage: "sin frequency",
		},

		cli.IntFlag{
			Name:  "rate, r",
			Value: 48000,
			Usage: "sampling rate",
		},
	}
	app.Action = generate
	app.Run(os.Args)
}
