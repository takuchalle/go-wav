package main

import (
	"fmt"
	"math"
	"os"

	"github.com/takuyaohashi/go-wav"
	"github.com/urfave/cli"
)

func generateAction(c *cli.Context) error {
	outputFile := c.String("output")
	rate := c.Int("rate")
	amp := c.Float64("amp")
	freq := c.Int("f")
	sec := c.Float64("sec")

	fmt.Printf("Output file: %s\n", outputFile)
	fmt.Printf("sampling rate: %d\n", rate)
	fmt.Printf("amplitude: %f\n", amp)
	fmt.Printf("freq: %d Hz\n", freq)
	fmt.Printf("time: %f sec\n", sec)

	f, err := os.Create(outputFile)
	if err != nil {
		return cli.NewExitError(err, 1)
	}
	defer f.Close()

	var param = wav.WriterParam{SampleRate: uint32(rate), BitsPerSample: 16, NumChannels: 1, AudioFormat: wav.AudioFormatPCM}

	w, err := wav.NewWriter(f, param)
	if err != nil {
		return cli.NewExitError(err, 1)
	}
	defer w.Close()

	length := sec * float64(rate)
	delta := float64(freq) / float64(rate)
	pi2 := math.Pi * 2.0

	for i := 0; i < int(length); i += 256 {
		samples := make([]int16, 256)
		for j := 0; j < 256; j++ {
			data := amp * math.Sin(pi2*float64(i+j)*delta)
			if data > 1.0 {
				data = 1.0
			}
			if data < -1.0 {
				data = -1.0
			}
			samples[j] = int16(data * float64(math.MaxInt16))
		}
		w.WriteSamples(samples)
	}

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

		cli.Float64Flag{
			Name:  "amp, a",
			Value: 1.0,
			Usage: "amplitude",
		},

		cli.IntFlag{
			Name:  "freq, f",
			Value: 1000,
			Usage: "sin frequency",
		},

		cli.IntFlag{
			Name:  "rate, r",
			Value: 48000,
			Usage: "sampling rate",
		},
	}
	app.Action = generateAction
	app.Run(os.Args)
}
