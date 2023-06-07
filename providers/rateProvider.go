package providers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type RateValue float32

type RateProvider struct {
}

func (provider *RateProvider) GetRate() (RateValue, error) {
	resp, err := http.Get("https://api.coingecko.com/api/v3/simple/price?ids=bitcoin&vs_currencies=uah")
	if err != nil {
		log.Println(err.Error())
		return 0, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	respJson := make(map[string]map[string]RateValue)
	json.Unmarshal(respBody, &respJson)

	return respJson["bitcoin"]["uah"], nil
}
