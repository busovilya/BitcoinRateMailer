package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/busovilya/CryptoRateMailer/models"
	"github.com/busovilya/CryptoRateMailer/services"
	"github.com/busovilya/CryptoRateMailer/types"
)

type SubscriptionHandler struct {
	subscriptionSvc services.SubscriptionService
}

func CreateSubscriptionHandler(subSvc services.SubscriptionService) *SubscriptionHandler {
	return &SubscriptionHandler{
		subscriptionSvc: subSvc,
	}
}

func (subscHandler *SubscriptionHandler) SubscribeHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err.Error())
		return
	}

	email := r.FormValue("email")
	coin := r.FormValue("coin")
	currency := r.FormValue("currency")

	err = subscHandler.subscriptionSvc.Subscribe(models.Subscription{
		Email:      email,
		Coin:       types.Coin(coin),
		VsCurrency: types.Currency(currency),
	})
	if err != nil {
		if err == services.SubscriptionExistsError {
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(err.Error())
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
		log.Println(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (subscHandler *SubscriptionHandler) SendEmailsHandler(w http.ResponseWriter, r *http.Request) {
	err := subscHandler.subscriptionSvc.SendEmails()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}
