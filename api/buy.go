package handler

import (
	"fmt"
	"net/http"
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

var bv bitvavo.Bitvavo
var conf *Config

func init() {
	conf = &Config{}
	multiconfig.New().MustLoad(conf)
	//make vercel copy-paste foolproof
	conf.ApiKey = strings.TrimSuffix(conf.ApiKey, "\n")
	conf.ApiSecret = strings.TrimSuffix(conf.ApiSecret, "\n")
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
	stringAmt := fmt.Sprintf("%.2f", eurAmount)
	return bv.PlaceOrder("BTC-EUR", "buy", "market", map[string]string{"amountQuote": stringAmt})
}

func Handler(w http.ResponseWriter, r *http.Request) {
	euros, err := GetEurBalance()
	if err != nil {
		logrus.WithError(err).Error("Something went wrong fetching Bitvavo account balance")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if euros == 0 {
		logrus.Info("No euros to buy coins")
		w.WriteHeader(http.StatusOK)
		return
	}
	order, err := BuyBitcoin(euros)
	if err != nil {
		logrus.WithError(err).Error("Something went wrong buying bitcoin on bitvavo")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	logrus.WithField("order", order).Info("Stacked some sats")
	w.WriteHeader(http.StatusOK)
}
