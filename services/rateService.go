package services

import (
	"errors"
	"log"

	"github.com/busovilya/BitcoinRateMailer/models"
	"github.com/busovilya/BitcoinRateMailer/providers"
	"github.com/busovilya/BitcoinRateMailer/types"
)

var CoinNotSupportedError = errors.New("Coin is not supported")

type RateService struct {
	rateProvider  providers.RateProvider
	coinsProvider providers.CoinsProvider
}

func CreateRateService(rateProvider providers.RateProvider, coinsProvider providers.CoinsProvider) *RateService {
	return &RateService{
		rateProvider:  rateProvider,
		coinsProvider: coinsProvider,
	}
}

func (rateSvc *RateService) GetRate(coin types.Coin, currency types.Currency) (*types.RateValue, error) {
	coins, err := rateSvc.coinsProvider.GetCoins()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	coinSupported := isCoinSupported(&coins, string(coin))
	if !coinSupported {
		return nil, CoinNotSupportedError
	}

	rate, err := rateSvc.rateProvider.GetRate(coin, currency)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return &rate, nil
}

func isCoinSupported(coins *[]models.Coin, coinId string) bool {
	for _, item := range *coins {
		if item.Id == coinId {
			return true
		}
	}

	return false
}
