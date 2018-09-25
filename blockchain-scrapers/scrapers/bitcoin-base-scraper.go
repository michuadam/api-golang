package blockchainscrapers

import (
	"github.com/diadata-org/api-golang/pkg/dia"
	"github.com/diadata-org/go-bitcoind"
	"log"
	"time"
)

const (
	SLEEP_TIME = 10
)

func RunScraper(host string, port int, user, password, symbol string, elapsedTime int) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	config := dia.GetConfigApi()
	if config == nil {
		log.Fatal("Couldn't load config")
	}

	client := dia.NewClient(config)
	if client == nil {
		log.Fatal("Couldn't load client")
	}

	bitcoinLib, err := bitcoind.New(host, port, user, password, USESSL)
	if err != nil {
		log.Fatal(err)
	}

	info, err := bitcoinLib.GetBlockchainInfo()
	if err != nil {
		log.Fatal(err)
	}

	for {
		medianTime := time.Unix(info.Mediantime, 0)
		blockTime := time.Now().Sub(medianTime)

		if blockTime < time.Duration(elapsedTime)*time.Second {
			txOutSetInfo, err := bitcoinLib.GetTxOutsetInfo()

			err = client.SendSupply(&dia.Supply{
				Symbol: symbol,
				CirculatingSupply: txOutSetInfo.TotalAmount,
			})
			if err != nil {
				log.Println("Error sending supply to API: ", err)
			} else {
				log.Println("Sent supply", txOutSetInfo.TotalAmount, "to API")
			}
		} else {
			log.Println("Sleeping until", symbol, "node is fully synchronized")
		}

		time.Sleep(SLEEP_TIME *time.Second)
	}
}
