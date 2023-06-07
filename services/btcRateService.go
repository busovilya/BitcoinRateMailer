package services

import (
	"log"

	"github.com/busovilya/BitcoinRateMailer/providers"
	viewmodels "github.com/busovilya/BitcoinRateMailer/viewModels"
)

type RateService struct {
	rateProvider providers.RateProvider
}

func CreateRateService(rateProvider providers.RateProvider) *RateService {
	return &RateService{
		rateProvider: rateProvider,
	}
}

func (rateSvc *RateService) GetBtcUahRate() (*viewmodels.RateView, error) {
	rate, err := rateSvc.rateProvider.GetRate()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return &viewmodels.RateView{Price: rate}, nil
}
