package viewmodels

import "github.com/busovilya/BitcoinRateMailer/providers"

type RateView struct {
	Price providers.RateValue `json:"price"`
}
