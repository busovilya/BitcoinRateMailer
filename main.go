package main

import (
	"github.com/busovilya/BitcoinRateMailer/handlers"
 	"net/http"
	"log"
)

func main() {
	http.HandleFunc("/rate", handlers.RateHandler)	
	http.HandleFunc("/subscribe", handlers.SubscribeHandler)	
	http.HandleFunc("/sendEmails", handlers.SendEmailsHandler)
	log.Fatal(http.ListenAndServe(":10000", nil))
}