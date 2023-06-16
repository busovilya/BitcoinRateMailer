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

	"github.com/busovilya/CryptoRateMailer/models"
	"github.com/busovilya/CryptoRateMailer/repositories"
)

var SubscriptionExistsError = errors.New("Subscription already exists")

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
	isValid := ValidateSubscription(&subscription)
	if !isValid {
		return errors.New(fmt.Sprintf("%s is not valid subscription", subscription.String()))
	}

	emailExists, err := subscService.subscriptionRepo.SubscriptionExists(subscription)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	if emailExists {
		return SubscriptionExistsError
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

	messages := make(map[string]string)
	for _, subscr := range subscriptions {
		rate, err := subscService.rateSvc.GetRate(subscr.Coin, subscr.VsCurrency)
		if err != nil {
			log.Println(err.Error())
			return err
		}

		msg, ok := messages[subscr.Email]
		if !ok {
			msg = ""
		}
		messages[subscr.Email] = fmt.Sprintf(
			"%s\n%s/%s: %s\n",
			msg,
			string(subscr.Coin),
			string(subscr.VsCurrency),
			strconv.FormatFloat(float64(*rate), 'f', -1, 64))

	}

	for email, msg := range messages {
		err = SendEmail(email, msg, "Crypto rates")
		if err != nil {
			log.Println(err.Error())
			return err
		}
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

func SendEmail(email string, msg string, subject string) error {
	from := os.Getenv("ENV_SENDER_EMAIL")
	password := os.Getenv("ENV_SENDER_PASSWORD")
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	message := []byte(fmt.Sprintf("Subject:%s\nTo:%s\n\n%s", subject, email, msg))
	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{email}, message)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func ValidateSubscription(subscr *models.Subscription) bool {
	return ValidateEmail(subscr.Email) &&
		len(subscr.Coin) != 0 &&
		len(subscr.VsCurrency) != 0
}

func ValidateEmail(email string) bool {
	_, err := mail.ParseAddress(email)

	if err != nil {
		log.Println(err.Error())
		return false
	}

	return true
}
