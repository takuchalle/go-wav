package wavplotter

import (
	"fmt"
	"math"

	"github.com/takuyaohashi/go-wav"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

type WavPlotter struct {
	w *wav.Reader
}

// Plot ...
func NewPlotter(wav *wav.Reader) (plotter *WavPlotter) {
	plotter = &WavPlotter{}
	plotter.w = wav
	return plotter
}

func (wp *WavPlotter) getTitleText() (name string) {
	return fmt.Sprintf("%dch\n%dHz", wp.w.GetNumChannels(), wp.w.GetSampleRate())
}

// pontSamples  ...
func (wp *WavPlotter) pointSamples(samples []int16) plotter.XYs {
	pts := make(plotter.XYs, len(samples))
	for i := range pts {
		pts[i].X = float64(i)
		pts[i].Y = (float64)(samples[i]) / float64(math.MaxInt16)
	}
	return pts
}

func (wp *WavPlotter) Output(filename string) {
	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	//label
	p.Title.Text = wp.getTitleText()

	// 座標範囲
	p.X.Min = 0
	p.Y.Min = -1
	p.Y.Max = 1

	data, err := wp.w.ReadSamples(1024)
	if err != nil {
		panic(err)
	}
	var pts plotter.XYs
	switch v := data.(type) {
	case []int16:
		pts = wp.pointSamples(v)
	default:
		panic(err)
	}

	plot1, err := plotter.NewScatter(pts)
	if err != nil {
		panic(err)
	}

	p.Add(plot1)

	if err := p.Save(8*vg.Inch, 6*vg.Inch, filename); err != nil {
		panic(err)
	}
}
