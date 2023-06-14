package providers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/busovilya/BitcoinRateMailer/models"
	"github.com/stretchr/testify/assert"
)

func TestGetCurrencies_EmptyList(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("[]"))
		return
	}))
	defer server.Close()

	provider := CreateRestAPICurrencyProvider(server.URL)
	result, err := provider.GetSupportedCurrencies()

	if err != nil {
		t.Errorf(err.Error())
	}

	assert.Zero(t, len(result))
}

func TestGetSupportedCurrencies_NonEmptyList(t *testing.T) {
	expected := []string{"currency1", "currency2"}
	expectedJson, _ := json.Marshal(expected)

	var expectedCurrencies []models.Currency
	json.Unmarshal(expectedJson, &expectedCurrencies)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expected)
		return
	}))
	defer server.Close()

	provider := CreateRestAPICurrencyProvider(server.URL)
	result, err := provider.GetSupportedCurrencies()

	if err != nil {
		t.Errorf(err.Error())
	}

	assert.Equal(t, expectedCurrencies, result)
}

func TestGetCurrencies_StatusNotOk(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}))
	defer server.Close()

	provider := CreateRestAPICurrencyProvider(server.URL)
	result, err := provider.GetSupportedCurrencies()

	assert.Nil(t, result)
	if assert.Error(t, err) {
		assert.Equal(t, err, APIResponseNotOKError)
	}
}
