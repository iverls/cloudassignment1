package handler

import (
	"encoding/json"
	"github.com/iverls/assignment1_country_info/fetch"
	"net/http"
	"strconv"
	"strings"
)

func CountryInfoHandler(w http.ResponseWriter, r *http.Request) {
	code := strings.TrimPrefix(r.URL.Path, ApiBase+InfoPath)

	result, err := fetch.GetCountryInfo(code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	limitParam := r.URL.Query().Get("limit")
	limit, _ := strconv.Atoi(limitParam) // defaults to 0 if empty

	cities, err := fetch.GetCities(result.Name, limit)
	if err == nil {
		result.Cities = cities
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
