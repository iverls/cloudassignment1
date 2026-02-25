// Handles the /info/{code} endpoint.
// Fetches general information about a country and formats it according to our spec.
package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/iverls/assignment1_country_info/fetch"
	"github.com/iverls/assignment1_country_info/model"
)

func CountryInfoHandler(w http.ResponseWriter, r *http.Request) {
	code := strings.TrimPrefix(r.URL.Path, ApiBase+InfoPath)

	rawCountry, err := fetch.GetCountryData(code)
	if err != nil {
		http.Error(w, "Country not found", http.StatusNotFound)
		return
	}

	capital := ""
	if len(rawCountry.Capital) > 0 {
		capital = rawCountry.Capital[0]
	}

	result := model.CountryInfoResponse{
		Name:       rawCountry.Name.Common,
		Continents: rawCountry.Continents,
		Population: rawCountry.Population,
		Area:       rawCountry.Area,
		Languages:  rawCountry.Languages,
		Borders:    rawCountry.Borders,
		Flag:       rawCountry.Flags.Png,
		Capital:    capital,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
