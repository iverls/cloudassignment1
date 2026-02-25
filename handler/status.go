package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/iverls/assignment1_country_info/fetch"
)

var startTime = time.Now()

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	status := fetch.CheckApiStatus()
	status.Version = "v1"
	status.Uptime = time.Since(startTime).Seconds()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(status)
}
