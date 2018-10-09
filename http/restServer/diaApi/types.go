package diaApi

import (
	"github.com/diadata-org/api-golang/internal/pkg/model"
	"github.com/diadata-org/api-golang/pkg/dia"
	"time"
)

type Coin struct {
	Symbol             string
	Name               string
	Price              float64
	PriceYesterday     *float64
	VolumeYesterdayUSD float64
	CirculatingSupply  float64
	Time               time.Time
}

type Coins struct {
	Coins []Coin
}
