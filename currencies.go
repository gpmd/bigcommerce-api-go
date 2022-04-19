package bigcommerce

import (
	"encoding/json"
	"log"
	"net/http"
)

// Currency is entry for BC currency API
type Currency struct {
	ID                     int      `json:"id"`
	IsDefault              bool     `json:"is_default"`
	LastUpdated            string   `json:"last_updated"`
	CountryIso2            string   `json:"country_iso2"`
	DefaultForCountryCodes []string `json:"default_for_country_codes"`
	CurrencyCode           string   `json:"currency_code"`
	CurrencyExchangeRate   string   `json:"currency_exchange_rate"`
	Name                   string   `json:"name"`
	Token                  string   `json:"token"`
	AutoUpdate             bool     `json:"auto_update"`
	TokenLocation          string   `json:"token_location"`
	DecimalToken           string   `json:"decimal_token"`
	ThousandsToken         string   `json:"thousands_token"`
	DecimalPlaces          int      `json:"decimal_places"`
	Enabled                bool     `json:"enabled"`
	IsTransactional        bool     `json:"is_transactional"`
	UseDefaultName         bool     `json:"use_default_name"`
}

// GetCurrencies returns the store's defined currencies
func (bc *Client) GetCurrencies() ([]Currency, error) {
	url := "/v2/currencies"

	req := bc.getAPIRequest(http.MethodGet, url, nil)
	res, err := bc.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := processBody(res)
	if err != nil {
		return nil, err
	}

	var cs []Currency
	err = json.Unmarshal(body, &cs)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return cs, nil
}
