package main

import (
	"github.com/diadata-org/api-golang/dia/helpers/configCollectors"
	"github.com/diadata-org/api-golang/frontend"
	"github.com/diadata-org/api-golang/services/model"
	"log"
	"os"
	"time"
)

const (
	GRAPH_PATH 	= "/charts/"
)


func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	dataStore, err := models.NewDataStore()
	if err != nil {
		log.Println(err)
		return
	}

	configCollector := configCollectors.NewConfigCollectors()
	if configCollector == nil {
		log.Println("Couldn't get config collector")
		return
	}

	for {
		for _, symbol := range configCollector.AllSymbols() {
			points, err := dataStore.GetChartPoints(symbol)
			if err != nil {
				log.Println(err)
			}

			if len(points) <= 0 {
				//log.Println("No datapoints for symbol", symbol)
				continue
			}

			log.Println("Producing chart for", symbol, "with", len(points),"datapoints")
			timePoints := make([]int64, len(points))
			pricePoints := make([]float64, len(points))

			for i := 0; i < len(points); i++ {
				timePoints[i] = points[i].UnixTime
				pricePoints[i] = points[i].Value
			}

			if _, err := os.Stat(GRAPH_PATH); os.IsNotExist(err) {
				err = os.MkdirAll(GRAPH_PATH, os.ModeDir | os.ModePerm)
				log.Println(err)
			}

			err = frontend.PriceGraph(pricePoints, timePoints, GRAPH_PATH+symbol+".png")
			if err != nil {
				log.Println(err)
			} else {
				log.Println("Created graph for"+symbol)
			}

		}

		time.Sleep(time.Minute*2)
	}
}
