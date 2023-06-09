package providers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/busovilya/BitcoinRateMailer/types"
)

type RateProvider struct {
}

func (provider *RateProvider) GetRate(token types.Coin, currency types.Currency) (types.RateValue, error) {
	url := fmt.Sprintf(
		"https://api.coingecko.com/api/v3/simple/price?ids=%s&vs_currencies=%s",
		token, currency)
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err.Error())
		return 0, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	respJson := make(map[types.Coin]map[types.Currency]types.RateValue)
	json.Unmarshal(respBody, &respJson)

	return respJson[token][currency], nil
}
