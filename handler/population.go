package handler

import (
	"encoding/json"
	"github.com/iverls/assignment1_country_info/fetch"
	"net/http"
	"strings"
)

func PopulationHandler(w http.ResponseWriter, r *http.Request) {
	code := strings.TrimPrefix(r.URL.Path, ApiBase+PopulationPath)
	limit := r.URL.Query().Get("limit")
	result, err := fetch.GetPopulationInfo(code, limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
