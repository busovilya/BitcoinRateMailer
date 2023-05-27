package services

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

func GetBtcUahRate() (int, error) {
	resp, err := http.Get("https://api.coingecko.com/api/v3/simple/price?ids=bitcoin&vs_currencies=uah")
	if err != nil {
		fmt.Println(err.Error())
		return -1, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	respJson := make(map[string]map[string]int)
	json.Unmarshal(respBody, &respJson)

	return respJson["bitcoin"]["uah"], nil
}
