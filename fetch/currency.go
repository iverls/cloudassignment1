package fetch

import "github.com/iverls/assignment1_country_info/model"

// GetExchangeRates fetches exchange rates for a specific base currency.
func GetExchangeRates(baseCurrency string) (model.CurrencyApiResponse, error) {
	var data model.CurrencyApiResponse
	err := FetchJSON("http://129.241.150.113:9090/currency/"+baseCurrency, &data)
	if err != nil {
		return model.CurrencyApiResponse{}, err
	}
	return data, nil
}
