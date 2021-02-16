package handler

import (
	"fmt"
	"strconv"

	"github.com/kiwiidb/go-bitvavo-api"
	"github.com/koding/multiconfig"
)

var bv bitvavo.Bitvavo
var conf *Config

func init() {
	conf = &Config{}
	multiconfig.New().MustLoad(conf)
	bv = bitvavo.Bitvavo{
		ApiKey:       conf.ApiKey,
		ApiSecret:    conf.ApiSecret,
		RestUrl:      conf.RestUrl,
		WsUrl:        conf.WsUrl,
		AccessWindow: 60000,
		WS:           bitvavo.Websocket{},
	}

}

func GetEurBalance() (balance float64, err error) {
	response, err := bv.Balance(map[string]string{"symbol": "EUR"})
	if err != nil {
		return 0, err
	}
	if len(response) != 1 {
		return 0, fmt.Errorf("Response balance should have length 1")
	}
	return strconv.ParseFloat(response[0].Available, 32)
}

func BuyBitcoin(eurAmount float64) (order bitvavo.Order, err error) {
	return bv.PlaceOrder("BTC-EUR", "buy", "market", map[string]string{"amountQuote": fmt.Sprintf("%f", eurAmount)})
}
