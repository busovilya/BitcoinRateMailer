package handlers

import (
	"log"
	"net/http"

	"github.com/busovilya/BitcoinRateMailer/models"
	"github.com/busovilya/BitcoinRateMailer/services"
	"github.com/busovilya/BitcoinRateMailer/types"
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
		w.WriteHeader(http.StatusConflict)
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
