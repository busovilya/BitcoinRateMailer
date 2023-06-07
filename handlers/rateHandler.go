package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/busovilya/BitcoinRateMailer/services"
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
	rate, err := rateHandler.rateSvc.GetBtcUahRate()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(rate)
}
