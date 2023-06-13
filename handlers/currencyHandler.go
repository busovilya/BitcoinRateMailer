package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/busovilya/BitcoinRateMailer/services"
)

type CurrencyHandler struct {
	currencySvc services.CurrencyService
}

func CreateCurrencyHandler(svc *services.CurrencyService) *CurrencyHandler {
	return &CurrencyHandler{currencySvc: *svc}
}

func (handler *CurrencyHandler) HandleCurrencies(w http.ResponseWriter, r *http.Request) {
	currencies, err := handler.currencySvc.GetSupportedCurrencies()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(currencies)
}
