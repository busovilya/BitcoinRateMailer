package providers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/busovilya/BitcoinRateMailer/types"
	"github.com/stretchr/testify/assert"
)

func TestGetRate_SuccesfullResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{\"coin\":{\"currency\":100}}"))
		return
	}))
	defer server.Close()

	provider := CreateRestAPIRateProvider(server.URL)
	result, err := provider.GetRate(types.Coin("coin"), types.Currency("currency"))

	if err != nil {
		t.Errorf(err.Error())
	}

	assert.Equal(t, result, types.RateValue(100))
}

func TestGetRate_WrongDataFormat(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{\"coin\":100"))
		return
	}))
	defer server.Close()

	provider := CreateRestAPIRateProvider(server.URL)
	_, err := provider.GetRate(types.Coin("coin"), types.Currency("currency"))

	assert.Error(t, err)
}

func TestGetRate_FailedRequest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}))
	defer server.Close()

	provider := CreateRestAPIRateProvider(server.URL)
	_, err := provider.GetRate(types.Coin("coin"), types.Currency("currency"))

	if assert.Error(t, err) {
		assert.Equal(t, err, APIResponseNotOKError)
	}
}
