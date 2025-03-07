package fetch

import (
	"encoding/json"
	"net/http"
	"strings"
)

func FetchJSON(url string, target interface{}) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return json.NewDecoder(res.Body).Decode(target)
}

func PostJSON(url string, requestBody interface{}, target interface{}) error {
	reqBody, _ := json.Marshal(requestBody)
	res, err := http.Post(url, "application/json", strings.NewReader(string(reqBody)))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return json.NewDecoder(res.Body).Decode(target)
}
