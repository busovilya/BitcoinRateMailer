package main

import (
	"log"
	"net/http"

	"github.com/busovilya/BitcoinRateMailer/handlers"
	"github.com/busovilya/BitcoinRateMailer/providers"
	"github.com/busovilya/BitcoinRateMailer/repositories"
	"github.com/busovilya/BitcoinRateMailer/services"
	"github.com/gorilla/mux"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	router := mux.NewRouter()

	coinsSvc := services.CreateCoinsService(providers.CoinsProvider{})
	coinsHandler := handlers.CreateCoinsHandler(coinsSvc)
	router.HandleFunc("/coins", coinsHandler.HandleCoins)

	currenciesSvc := services.CreateCurrencyService(&providers.CurrencyProvider{})
	currenciesHandler := handlers.CreateCurrencyHandler(currenciesSvc)
	router.HandleFunc("/currencies", currenciesHandler.HandleCurrencies)

	rateSvc := services.CreateRateService(providers.RateProvider{}, providers.CoinsProvider{}, providers.CurrencyProvider{})
	rateHandler := handlers.CreateRateHandler(rateSvc)
	router.HandleFunc("/rate/{coin}/{currency}", rateHandler.HandleRateRequest)

	subscriptionHandler := handlers.CreateSubscriptionHandler(
		*services.CreateSubscriptionService(
			repositories.CreateSubscriptionFileRepo("emails.data"),
			*rateSvc),
	)
	router.HandleFunc("/subscribe", subscriptionHandler.SubscribeHandler)
	router.HandleFunc("/sendEmails", subscriptionHandler.SendEmailsHandler)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":10000", nil))
}
