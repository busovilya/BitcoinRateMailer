package providers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/busovilya/BitcoinRateMailer/models"
)

type CoinsProvider struct {
}

func (provider *CoinsProvider) GetCoins() ([]models.Coin, error) {
	url := fmt.Sprintf(
		"https://api.coingecko.com/api/v3/coins/list",
	)
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Failed HTTP request to API")
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
