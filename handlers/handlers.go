package handlers

import (
	"github.com/busovilya/BitcoinRateMailer/services"
	"fmt"
 	"net/http"
	"log"
	"os"
	"encoding/json"
)

func RateHandler(w http.ResponseWriter, r *http.Request) {
	rate, err := services.GetBtcUahRate()
	
	respJson := make(map[string]interface{})

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		respJson["result"] = "failed to retrive BTC rate"
		json.NewEncoder(w).Encode(respJson)	
		return
	}

	w.WriteHeader(http.StatusOK)
	respJson["btcuah"] = rate
	json.NewEncoder(w).Encode(respJson)
}

func SubscribeHandler(w http.ResponseWriter, r *http.Request) {
	respJson := make(map[string]interface{})

	err := r.ParseForm()
	if err != nil {
		fmt.Println(err.Error())
	}
	
	email := r.PostForm["email"]
	if len(email) > 0 {
		file, err := os.OpenFile("emails.data", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
		defer file.Close()
		
		if err != nil {
			log.Println("Can't open emails.data to read the list of subscribers")
			respJson["result"] = "failed to subscribe email"
			w.WriteHeader(http.StatusInternalServerError)
		} else if !services.FindEmailInFile(email[0], file) {
			services.WriteEmailToFile(email[0], file)
			
			respJson["result"] = "email added"
			w.WriteHeader(http.StatusOK)
		} else {
			respJson["result"] = "email exists"
			w.WriteHeader(http.StatusConflict)
		}
	} else {
			respJson["result"] = "email address is not provided"
			w.WriteHeader(http.StatusBadRequest)
	}
	  
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respJson)
}

func SendEmailsHandler(w http.ResponseWriter, r *http.Request) {
	respJson := make(map[string]interface{})
	w.Header().Set("Content-Type", "application/json")

	file, err := os.OpenFile("emails.data", os.O_RDONLY, 0644)
	defer file.Close()
	if err != nil {
		log.Println("Failed to send emails")
		w.WriteHeader(http.StatusInternalServerError)
		respJson["result"] = "can't send emails"
		json.NewEncoder(w).Encode(respJson)
		return
	}
	emails := services.ReadEmailsList(file)

	rate, err := services.GetBtcUahRate()	
	if err != nil {
		log.Println("Failed to retrieve BTC rate")
		w.WriteHeader(http.StatusInternalServerError)
		respJson["result"] = "can't send emails"
		json.NewEncoder(w).Encode(respJson)
		return
	}

	mailText := fmt.Sprintf("BTC/UAH: %d", rate)	
	for _, email := range emails {
		go services.SendEmail(email, mailText)
	}
	
	respJson["result"] = "emails were sent"
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(respJson)
}
