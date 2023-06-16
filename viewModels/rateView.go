package viewmodels

import (
	"github.com/busovilya/CryptoRateMailer/types"
)

type RateView struct {
	Token    types.Coin      `json:"coin"`
	Currency types.Currency  `json:"currency"`
	Price    types.RateValue `json:"price"`
}
