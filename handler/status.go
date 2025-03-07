package handler

import (
	"encoding/json"
	"github.com/iverls/assignment1_country_info/fetch"
	"net/http"
	"time"
)

var startTime = time.Now()

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	status := fetch.CheckApiStatus()
	status.Version = "v1"
	status.Uptime = time.Since(startTime).Seconds()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(status)
}
