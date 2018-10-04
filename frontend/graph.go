package frontend

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"image/color"
	"log"
	"math"
	"math/rand"
	"time"
)

const (
	IMAGE_HEIGHT = 2
	IMAGE_WIDTH = 6
)

func PriceGraph(prices []float64, times []int64, path string) error {
	graph, err := plot.New()
	if err != nil {
		return err
	}

	// normalize data
	dataPoints := make(plotter.XYs, len(prices))
	for i := range dataPoints {
		dataPoints[i].X = float64(times[i])/float64(len(times))
		dataPoints[i].Y = prices[i]
	}

	// insert data in line
	line, err := plotter.NewLine(dataPoints)
	if err != nil {
		return err
	}

	// change presentation
	graph.HideAxes()
	graph.BackgroundColor = color.Transparent
	line.LineStyle.Color = color.RGBA{R: 63, G: 63, B: 237, A: 255}
	//*line.ShadeColor = &color.RGBA{R: 63, G: 63, B: 237, A: 150}

	// add data to graph
	err = plotutil.AddLines(graph, line)
	if err != nil {
		return err
	}

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
	for i := range prices {
		prices[i] = math.Sin(float64(x))
		x += 0.01
	}

	weekAgo := time.Now().Add(-time.Hour*7*24)
	for i := range times {
		times[i] = weekAgo.Add(time.Minute * time.Duration(2*i)).Unix()
	}

	err := PriceGraph(prices, times, "example")
	if err != nil {
		log.Println(err)
	}
}