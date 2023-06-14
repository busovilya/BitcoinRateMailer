package providers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/busovilya/BitcoinRateMailer/models"
)

var APIResponseNotOKError = errors.New("Failed HTTP request to API")

type CoinsProvider interface {
	GetCoins() ([]models.Coin, error)
}

type RestAPICoinsProvider struct {
	url string
}

func CreateRestAPICoinsProvider(url string) CoinsProvider {
	return RestAPICoinsProvider{
		url: url,
	}
}

func CreateCoingeckoProviderCoins() CoinsProvider {
	return CreateRestAPICoinsProvider("https://api.coingecko.com/api/v3/coins/list")
}

func (provider RestAPICoinsProvider) GetCoins() ([]models.Coin, error) {
	reuqest, err := http.NewRequest(http.MethodGet, provider.url, nil)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	client := http.Client{}
	resp, err := client.Do(reuqest)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, APIResponseNotOKError
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var coinsList []models.Coin
	err = json.Unmarshal(body, &coinsList)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return coinsList, nil
}
