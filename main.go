package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kiwiidb/go-bitvavo-api"
	"github.com/koding/multiconfig"
	"github.com/sirupsen/logrus"
)

type Config struct {
	ApiKey    string `required:"true"`
	ApiSecret string `required:"true"`
	RestUrl   string `default:"https://api.bitvavo.com/v2"`
	WsUrl     string `default:"wss://ws.bitvavo.com/v2/"`
}

func main() {
	conf := &Config{}
	multiconfig.New().MustLoad(conf)
	//make copy-paste foolproof
	conf.ApiKey = strings.TrimSuffix(conf.ApiKey, "\n")
	conf.ApiSecret = strings.TrimSuffix(conf.ApiSecret, "\n")
	bv := bitvavo.Bitvavo{
		ApiKey:       conf.ApiKey,
		ApiSecret:    conf.ApiSecret,
		RestUrl:      conf.RestUrl,
		WsUrl:        conf.WsUrl,
		AccessWindow: 60000,
		WS:           bitvavo.Websocket{},
	}
	euros, err := GetEurBalance(&bv)
	if err != nil {
		logrus.WithError(err).Error("Something went wrong fetching euro balance")
		return
	}
	if euros == 0 {
		logrus.Debug("No euros to buy coins")
		return
	}
	order, err := BuyBitcoin(&bv, euros)
	if err != nil {
		logrus.WithError(err).Error("Something went wrong buying bitcoin on bitvavo")
		return
	}
	logrus.WithField("order", order).Debug("Stacked some sats")
}

func GetEurBalance(bv *bitvavo.Bitvavo) (balance float64, err error) {
	response, err := bv.Balance(map[string]string{"symbol": "EUR"})
	if err != nil {
		return 0, err
	}
	if len(response) != 1 {
		return 0, fmt.Errorf("Response balance should have length 1")
	}
	return strconv.ParseFloat(response[0].Available, 32)
}

func BuyBitcoin(bv *bitvavo.Bitvavo, eurAmount float64) (order bitvavo.Order, err error) {
	stringAmt := fmt.Sprintf("%.2f", eurAmount)
	return bv.PlaceOrder("BTC-EUR", "buy", "market", map[string]string{"amountQuote": stringAmt})
}
