package repositories

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/busovilya/BitcoinRateMailer/models"
)

type SubscriptionRepository interface {
	Add(*models.Subscription) error
	GetSubscriptions() ([]models.Subscription, error)
	SubscriptionExists(models.Subscription) (bool, error)
}

type SubscriptionFileRepo struct {
	file string
}

func CreateSubscriptionFileRepo(fileName string) *SubscriptionFileRepo {
	return &SubscriptionFileRepo{
		file: fileName,
	}
}

func (repo *SubscriptionFileRepo) Add(sub *models.Subscription) error {
	file, err := os.OpenFile(repo.file, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	defer file.Close()

	if err != nil {
		log.Println(err.Error())
		return err
	}

	_, err = file.WriteString(fmt.Sprintf("%s\n", sub.String()))
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func (repo *SubscriptionFileRepo) GetSubscriptions() ([]models.Subscription, error) {
	file, err := os.OpenFile(repo.file, os.O_RDONLY, 0644)
	defer file.Close()

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	var subscriptions []models.Subscription
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		subscr, err := models.ParseSubscription(line)
		if err != nil {
			return nil, err
		}

		subscriptions = append(subscriptions, *subscr)
	}

	return subscriptions, nil
}

func (repo *SubscriptionFileRepo) SubscriptionExists(subscription models.Subscription) (bool, error) {
	subs, err := repo.GetSubscriptions()
	if err != nil {
		log.Println(err.Error())
		return false, err
	}

	for _, item := range subs {
		if item == subscription {
			return true, nil
		}
	}

	return false, nil
}
