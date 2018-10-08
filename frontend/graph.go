package frontend

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"image/color"
	"log"
	"math"
	"math/rand"
	"time"
)

const (
	IMAGE_HEIGHT = 1
	IMAGE_WIDTH = 6
)

func PriceGraph(prices []float64, times []int64, path string) error {
	graph, err := plot.New()
	if err != nil {
		return err
	}

	// downsample data
	dataPoints := make(plotter.XYs, len(times))
	for i := 0; i < len(times); i++ {
		dataPoints[i].X = float64(times[i])
		dataPoints[i].Y = prices[i]
	}

	// insert data in line
	line, err := plotter.NewLine(dataPoints)
	if err != nil {
		return err
	}

	// change presentation
	graph.HideAxes()
	graph.BackgroundColor = color.RGBA{R: 0, G: 0, B: 0, A: 0}
	line.LineStyle.Color = color.RGBA{R: 0, G: 0, B: 255, A: 255}

	//*line.ShadeColor = color.RGBA{R:200, G: 220,B: 255, A: 250}

	// add line to graph
	graph.Add(line)

	// Save graph
	err = graph.Save(IMAGE_WIDTH*vg.Centimeter, IMAGE_HEIGHT*vg.Centimeter, path)
	if err != nil {
		return err
	}

	return nil
}

func test() {
	numPoints := 30*24*7
	prices := make([]float64, numPoints)
	times := make([]int64, numPoints)

	rand.Seed(time.Now().Unix())
	var x float64 = 0.0
	for i := 0; i < len(prices); i++ {
		even := rand.Uint32() % 2 == 0
		prices[i] = math.Sin(float64(rand.Float64()))
		if !even {
			prices[i] *= -1
		}
		x += 0.1
	}

	weekAgo := time.Now().Add(-time.Hour*7*24)
	for i := range times {
		times[i] = weekAgo.Add(time.Minute * time.Duration(2*i)).Unix()
	}

	err := PriceGraph(prices, times, "example.png")
	if err != nil {
		log.Println(err)
	}
}