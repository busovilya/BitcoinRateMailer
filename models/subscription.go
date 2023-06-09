package models

import (
	"errors"
	"fmt"
	"strings"

	"github.com/busovilya/BitcoinRateMailer/types"
)

type Subscription struct {
	Email      string
	Coin       types.Coin
	VsCurrency types.Currency
}

func (s *Subscription) String() string {
	return fmt.Sprintf("%s %s %s", s.Email, s.Coin, s.VsCurrency)
}

func ParseSubscription(str string) (*Subscription, error) {
	strParts := strings.Split(str, " ")

	if len(strParts) != 3 {
		return nil, errors.New("Subscription parsing failed. Wrong format")
	}

	return &Subscription{
		Email:      strParts[0],
		Coin:       types.Coin(strParts[1]),
		VsCurrency: types.Currency(strParts[2]),
	}, nil
}
