package providers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/busovilya/CryptoRateMailer/models"
)

type CurrencyProvider interface {
	GetSupportedCurrencies() ([]models.Currency, error)
}

type RestAPICurrencyProvider struct {
	url string
}

func CreateRestAPICurrencyProvider(url string) CurrencyProvider {
	return RestAPICurrencyProvider{
		url: url,
	}
}

func CreateCoingeckoCurrencyProvider() CurrencyProvider {
	return CreateRestAPICurrencyProvider("https://api.coingecko.com/api/v3/simple/supported_vs_currencies")
}

func (provider RestAPICurrencyProvider) GetSupportedCurrencies() ([]models.Currency, error) {
	resp, err := http.Get(provider.url)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, APIResponseNotOKError
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var currenciesList []models.Currency
	err = json.Unmarshal(body, &currenciesList)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return currenciesList, nil
}
