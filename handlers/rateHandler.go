package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/busovilya/CryptoRateMailer/services"
	"github.com/busovilya/CryptoRateMailer/types"
	viewmodels "github.com/busovilya/CryptoRateMailer/viewModels"
	"github.com/gorilla/mux"
)

type RateHandler struct {
	rateSvc services.RateService
}

func CreateRateHandler(rateSvc *services.RateService) *RateHandler {
	return &RateHandler{
		rateSvc: *rateSvc,
	}
}

func (rateHandler *RateHandler) HandleRateRequest(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	coin := vars["coin"]
	currency := vars["currency"]

	rate, err := rateHandler.rateSvc.GetRate(
		types.Coin(coin),
		types.Currency(currency))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if err == services.CoinNotSupportedError {
			json.NewEncoder(w).Encode(viewmodels.Error{Error: "coin is not supported"})
		}
		log.Println(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(rate)
}
