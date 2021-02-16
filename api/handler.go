package handler

import "net/http"

func handler(w http.ResponseWriter, r *http.Request) {
	GetEurBalance()
}
