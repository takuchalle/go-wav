package main

import (
	"fmt"
	"os"

	"github.com/takuyaohashi/go-wav"
	"github.com/takuyaohashi/go-wav/wavplotter"

	"github.com/urfave/cli"
)

// getFileName ...
func getFileName(givenName string, w *wav.Reader) string {
	if givenName == "wave.png" {
		return fmt.Sprintf("%dch_%dhz_wav.png",
			w.GetNumChannels(),
			w.GetSampleRate())
	} else {
		return givenName
	}
}

// plot ...
func plot(c *cli.Context) error {
	if c.Args().Get(0) == "" {
		return cli.NewExitError("Need input file", 1)
	}

	f, err := os.Open(c.Args().Get(0))
	if err != nil {
		return cli.NewExitError("No such file", 2)
	}
	defer f.Close()

	w, err := wav.NewReader(f)
	if err != nil {
		return cli.NewExitError(err, 3)
	}

	plotter := wavplotter.NewPlotter(w)

	filename := getFileName(c.String("output"), w)

	plotter.Output(filename)

	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "Wave Plotter"
	app.Usage = "Create Audio Wave Image file"
	app.Version = "0.0.1"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "output, o",
			Value: "wave.png",
			Usage: "output file name",
		},
	}
	app.Action = plot
	app.Run(os.Args)
}
