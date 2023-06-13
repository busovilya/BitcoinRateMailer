package services

import (
	"log"

	"github.com/busovilya/BitcoinRateMailer/providers"
	"github.com/busovilya/BitcoinRateMailer/types"
)

type CurrencyService struct {
	provider providers.CurrencyProvider
}

func CreateCurrencyService(provider *providers.CurrencyProvider) *CurrencyService {
	return &CurrencyService{
		provider: *provider,
	}
}

func (svc *CurrencyService) GetSupportedCurrencies() ([]types.Currency, error) {
	currencies, err := svc.provider.GetSupportedCurrencies()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	var currencyList []types.Currency
	for _, currency := range currencies {
		currencyList = append(currencyList, types.Currency(currency))
	}
	return currencyList, nil
}
