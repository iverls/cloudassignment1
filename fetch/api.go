package fetch

import (
	"encoding/json"
	"net/http"
)

func FetchJSON(url string, target interface{}) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return json.NewDecoder(res.Body).Decode(target)
}
