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

func (wp *WavPlotter) getFileName() (name string) {
	return fmt.Sprintf("%dch_wav.png", wp.w.GetNumChannels())
}

func (wp *WavPlotter) getTitleText() (name string) {
	return fmt.Sprintf("%dch \n", wp.w.GetNumChannels())
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

func (wp *WavPlotter) Plot() {
	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	//label
	p.Title.Text = wp.getTitleText()

	// 補助線
	p.Add(plotter.NewGrid())

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

	//plot1,plot2をplot
	p.Add(plot1)

	// plot.pngに保存
	if err := p.Save(4*vg.Inch, 4*vg.Inch, wp.getFileName()); err != nil {
		panic(err)
	}

	fmt.Printf("Saved %s\n", wp.getFileName())
}
