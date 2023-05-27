package services

import (
	"net/smtp"
	"fmt"
	"log"
	"os"
)

func SendEmail(email string, text string) {
	from := os.Getenv("ENV_SENDER_EMAIL")
	password := os.Getenv("ENV_SENDER_PASSWORD")
	to := email
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	message := []byte(fmt.Sprintf("Subject:BTC/UAH rate\nTo:%s\n\n%s", to, text))
	auth := smtp.PlainAuth("", from, password, smtpHost) 

	err := smtp.SendMail(smtpHost + ":" + smtpPort, auth, from, []string{to}, message)
 	if err != nil {
   		log.Println(err.Error())
    	return
  	}
  	log.Println("Email Sent Successfully!")
}
