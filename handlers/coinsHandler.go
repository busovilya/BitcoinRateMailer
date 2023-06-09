package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/busovilya/BitcoinRateMailer/services"
)

type CoinsHandler struct {
	coinsSvc services.CoinsService
}

func CreateCoinsHandler(coinsSvc *services.CoinsService) *CoinsHandler {
	return &CoinsHandler{
		coinsSvc: *coinsSvc,
	}
}

func (coinsHandler *CoinsHandler) CoinsHandler(w http.ResponseWriter, r *http.Request) {
	coins, err := coinsHandler.coinsSvc.GetCoins()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(coins)
}
