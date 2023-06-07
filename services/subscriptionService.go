package services

import (
	"errors"
	"fmt"
	"log"
	"net/mail"
	"net/smtp"
	"os"
	"strconv"
	"strings"

	"github.com/busovilya/BitcoinRateMailer/models"
	"github.com/busovilya/BitcoinRateMailer/repositories"
)

type SubscriptionService struct {
	subscriptionRepo repositories.SubscriptionRepository
	rateSvc          RateService
}

func CreateSubscriptionService(
	subRepo repositories.SubscriptionRepository,
	rateService RateService) *SubscriptionService {
	return &SubscriptionService{
		subscriptionRepo: subRepo,
		rateSvc:          rateService,
	}
}

func (subscService *SubscriptionService) Subscribe(subscription models.Subscription) error {
	isValid := ValidateEmail(subscription.Email)
	if !isValid {
		return errors.New(fmt.Sprintf("%s is not valid email", subscription.Email))
	}

	emailExists, err := subscService.subscriptionRepo.SubscriptionExists(subscription)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	if emailExists {
		return errors.New(fmt.Sprintf("%s already exists", subscription.Email))
	}

	err = subscService.subscriptionRepo.Add(&subscription)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func (subscService *SubscriptionService) SendEmails() error {
	subscriptions, err := subscService.subscriptionRepo.GetSubscriptions()
	if err != nil {
		log.Println(err.Error())
		return err
	}

	rate, err := subscService.rateSvc.GetBtcUahRate()
	if err != nil {
		log.Println(err.Error())
		return err
	}

	emails := []string{}
	for _, sub := range subscriptions {
		emails = append(emails, sub.Email)
	}

	err = SendEmails(emails, strconv.FormatFloat(float64(rate.Price), 'f', -1, 64), "BTC/UAH")
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func SendEmails(emails []string, msg string, subject string) error {
	from := os.Getenv("ENV_SENDER_EMAIL")
	password := os.Getenv("ENV_SENDER_PASSWORD")
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	message := []byte(fmt.Sprintf("Subject:%s\nTo:%s\n\n%s", subject, strings.Join(emails, ", "), msg))
	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, emails, message)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func ValidateEmail(email string) bool {
	_, err := mail.ParseAddress(email)

	if err != nil {
		log.Println(err.Error())
		return false
	}

	return true
}
