package providers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/busovilya/BitcoinRateMailer/types"
)

type RateProvider interface {
	GetRate(types.Coin, types.Currency) (types.RateValue, error)
}

type RestAPIRateProvider struct {
	url string
}

func CreateRestAPIRateProvider(url string) RestAPIRateProvider {
	return RestAPIRateProvider{
		url: url,
	}
}

func CreateCoingeckoRateProvider() RestAPIRateProvider {
	return CreateRestAPIRateProvider("https://api.coingecko.com/api/v3/simple/price")
}

func (provider RestAPIRateProvider) GetRate(token types.Coin, currency types.Currency) (types.RateValue, error) {
	url := fmt.Sprintf(
		"%s?ids=%s&vs_currencies=%s",
		provider.url, token, currency)
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}

	if resp.StatusCode != http.StatusOK {
		return 0, APIResponseNotOKError
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	respJson := make(map[types.Coin]map[types.Currency]types.RateValue)
	err = json.Unmarshal(respBody, &respJson)
	if err != nil {
		return 0, err
	}

	return respJson[token][currency], nil
}
