package providers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/busovilya/BitcoinRateMailer/models"
)

type CurrencyProvider struct{}

func (provider *CurrencyProvider) GetSupportedCurrencies() ([]models.Currency, error) {
	url := fmt.Sprintf(
		"https://api.coingecko.com/api/v3/simple/supported_vs_currencies",
	)
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err.Error())
		return nil, err
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
