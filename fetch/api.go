package fetch

import (
	"encoding/json"
	"net/http"
	"time"
)

// Custom HTTP client with 10 second timeout.
var client = &http.Client{
	Timeout: 10 * time.Second,
}

func FetchJSON(url string, target interface{}) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return json.NewDecoder(res.Body).Decode(target)
}
