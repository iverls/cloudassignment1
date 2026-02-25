// Handles the /exchange/{code} endpoint.
// Finds the given country's currency, looks up the currencies of all its
// neighboring countries, and returns the exchange rates between them.
package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/iverls/assignment1_country_info/fetch"
	"github.com/iverls/assignment1_country_info/model"
)

func ExchangeHandler(w http.ResponseWriter, r *http.Request) {
	code := strings.TrimPrefix(r.URL.Path, ApiBase+ExchangePath)

	baseCountry, err := fetch.GetCountryData(code)
	if err != nil {
		http.Error(w, "Base country not found", http.StatusNotFound)
		return
	}

	baseCurrency := ""
	for curCode := range baseCountry.Currencies {
		baseCurrency = curCode
		break
	}

	ratesData, err := fetch.GetExchangeRates(baseCurrency)
	if err != nil {
		http.Error(w, "Failed to fetch exchange rates", http.StatusInternalServerError)
		return
	}

	borderCurrencies := make(map[string]bool)
	for _, borderCode := range baseCountry.Borders {
		borderCountry, err := fetch.GetCountryData(borderCode)
		if err == nil {
			for curCode := range borderCountry.Currencies {
				borderCurrencies[curCode] = true
			}
		}
	}

	filteredRates := make(map[string]float64)
	for borderCurCode := range borderCurrencies {
		if rate, exists := ratesData.Rates[borderCurCode]; exists {
			filteredRates[borderCurCode] = rate
		}
	}

	result := model.ExchangeResponse{
		Country:       baseCountry.Name.Common,
		BaseCurrency:  baseCurrency,
		ExchangeRates: filteredRates,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
