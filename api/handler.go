package handler

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func handler(w http.ResponseWriter, r *http.Request) {
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
	logrus.WithField("order", order.Amount).Info("Stacked some sats")
	w.WriteHeader(http.StatusOK)
}
