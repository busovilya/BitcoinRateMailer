package services

import (
	"log"

	"github.com/busovilya/BitcoinRateMailer/providers"
	"github.com/busovilya/BitcoinRateMailer/types"
)

type CoinsService struct {
	coinsProvider providers.CoinsProvider
}

func CreateCoinsService(coinsProvider providers.CoinsProvider) *CoinsService {
	return &CoinsService{
		coinsProvider: coinsProvider,
	}
}

func (coinsSvc *CoinsService) GetCoins() ([]types.Coin, error) {
	coins, err := coinsSvc.coinsProvider.GetCoins()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	var coinsList []types.Coin
	for _, coin := range coins {
		coinsList = append(coinsList, types.Coin(coin.Id))
	}
	return coinsList, nil
}
