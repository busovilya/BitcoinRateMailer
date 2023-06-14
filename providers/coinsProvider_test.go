package providers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/busovilya/BitcoinRateMailer/models"
	"github.com/stretchr/testify/assert"
)

func TestGetCoins_EmptyList(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("[]"))
		return
	}))
	defer server.Close()

	provider := CreateRestAPICoinsProvider(server.URL)
	result, err := provider.GetCoins()

	if err != nil {
		t.Errorf(err.Error())
	}

	assert.Zero(t, len(result))
}

func TestGetCoins_NonEmptyList(t *testing.T) {
	expected := []map[string]string{{
		"id":   "coin1",
		"name": "coin1_name",
	}, {
		"id":   "coin2",
		"name": "coin2_name",
	}}
	expectedJson, _ := json.Marshal(expected)

	var expectedCoins []models.Coin
	json.Unmarshal(expectedJson, &expectedCoins)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expected)
		return
	}))
	defer server.Close()

	provider := CreateRestAPICoinsProvider(server.URL)
	result, err := provider.GetCoins()

	if err != nil {
		t.Errorf(err.Error())
	}

	assert.Equal(t, expectedCoins, result)
}

func TestGetCoins_DifferentFields(t *testing.T) {
	expected := []map[string]string{{
		"id":        "coin1",
		"name":      "coin1_name",
		"new_field": "data",
	}, {
		"id": "coin2",
	}}
	expectedJson, _ := json.Marshal(expected)

	var expectedCoins []models.Coin
	json.Unmarshal(expectedJson, &expectedCoins)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(expected)
		return
	}))
	defer server.Close()

	provider := CreateRestAPICoinsProvider(server.URL)
	result, err := provider.GetCoins()

	if err != nil {
		t.Errorf(err.Error())
	}

	assert.Equal(t, expectedCoins, result)
}

func TestGetCoins_StatusNotOk(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}))
	defer server.Close()

	provider := CreateRestAPICoinsProvider(server.URL)
	result, err := provider.GetCoins()

	assert.Nil(t, result)
	if assert.Error(t, err) {
		assert.Equal(t, err, APIResponseNotOKError)
	}
}
