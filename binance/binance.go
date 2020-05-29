package binance

import (
	"bgl/db"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/jasonlvhit/gocron"
)

const url = "https://api.binance.us/api/v3/ticker/price?symbol=BTCUSD"

type btcQuote struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

// GetBtcPriceEveryMinute is a background job fetching lastest price from remote server
func GetBtcPriceEveryMinute(db *db.DB) {
	// Fetch price when program starts
	GetPriceFromBinance(db)

	// Start periodical price fetching job in background
	gocron.Every(1).Minute().Do(GetPriceFromBinance, db)
	<-gocron.Start()
}

// GetPriceFromBinance fetch data from binance API and save data to DB
func GetPriceFromBinance(db *db.DB) {
	tick := time.Now().Unix()
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	var quote btcQuote
	err = json.Unmarshal(body, &quote)
	if err != nil {
		log.Println(err)
	}

	price, _ := strconv.ParseFloat(quote.Price, 64)

	err = db.InsertNewPriceQuote(price, tick)
	if err != nil {
		log.Println(err)
	}
}
