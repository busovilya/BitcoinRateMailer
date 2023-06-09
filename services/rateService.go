package services

import (
	"errors"
	"log"

	"github.com/busovilya/CryptoRateMailer/models"
	"github.com/busovilya/CryptoRateMailer/providers"
	"github.com/busovilya/CryptoRateMailer/types"
)

var CoinNotSupportedError = errors.New("Coin is not supported")
var CurrencyNotSupportedError = errors.New("Currency is not supported")

type RateService struct {
	rateProvider       providers.RateProvider
	coinsProvider      providers.CoinsProvider
	currenciesProvider providers.CurrencyProvider
}

func CreateRateService(
	rateProvider providers.RateProvider,
	coinsProvider providers.CoinsProvider,
	currenciesProvider providers.CurrencyProvider) *RateService {
	return &RateService{
		rateProvider:       rateProvider,
		coinsProvider:      coinsProvider,
		currenciesProvider: currenciesProvider,
	}
}

func (rateSvc *RateService) GetRate(coin types.Coin, currency types.Currency) (*types.RateValue, error) {
	coins, err := rateSvc.coinsProvider.GetCoins()
	if err != nil {
		return nil, err
	}

	coinSupported := isCoinSupported(&coins, string(coin))
	if !coinSupported {
		return nil, CoinNotSupportedError
	}

	currencies, err := rateSvc.currenciesProvider.GetSupportedCurrencies()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	currencySupported := isCurrencySupported(&currencies, models.Currency(currency))

	if !currencySupported {
		return nil, CurrencyNotSupportedError
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

func isCurrencySupported(currencies *[]models.Currency, currency models.Currency) bool {
	for _, item := range *currencies {
		if item == currency {
			return true
		}
	}

	return false
}
