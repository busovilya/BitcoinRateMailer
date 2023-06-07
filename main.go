package main

import (
	"log"
	"net/http"

	"github.com/busovilya/BitcoinRateMailer/handlers"
	"github.com/busovilya/BitcoinRateMailer/providers"
	"github.com/busovilya/BitcoinRateMailer/repositories"
	"github.com/busovilya/BitcoinRateMailer/services"
)

func main() {
	rateSvc := services.CreateRateService(providers.RateProvider{})
	rateHandler := handlers.CreateRateHandler(rateSvc)
	http.HandleFunc("/rate", rateHandler.HandleRateRequest)

	subscriptionHandler := handlers.CreateSubscriptionHandler(
		*services.CreateSubscriptionService(
			repositories.CreateSubscriptionFileRepo("emails.data"),
			*rateSvc),
	)
	http.HandleFunc("/subscribe", subscriptionHandler.SubscribeHandler)
	http.HandleFunc("/sendEmails", subscriptionHandler.SendEmailsHandler)
	log.Fatal(http.ListenAndServe(":10000", nil))
}
