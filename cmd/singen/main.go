package main

import (
	"os"

	"github.com/takuyaohashi/go-wav"
	"github.com/urfave/cli"
)

func generate(c *cli.Context) error {
	f, err := os.Create(c.String("output"))
	if err != nil {
		return cli.NewExitError(err, 1)
	}
	defer f.Close()

	w, err := wav.NewWriter(f, wav.WriterParam{SampleRate: 1, BitsPerSample: 1, NumChannels: 1, AudioFormat: wav.AudioFormatPCM})
	if err != nil {
		return cli.NewExitError(err, 1)
	}
	defer w.Close()

	w.WriteSamples()

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
	}
	app.Action = generate
	app.Run(os.Args)
}
