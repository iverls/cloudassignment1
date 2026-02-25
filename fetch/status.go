package fetch

import (
	"github.com/iverls/assignment1_country_info/model"
	"net/http"
)

func CheckApiStatus() model.StatusResponse {
	restCountriesResp, err1 := http.Get("http://129.241.150.113:8080/v3.1/all?fields=name")
	currenciesResp, err2 := http.Get("http://129.241.150.113:9090/currency/NOK")

	return model.StatusResponse{
		RestCountriesApi: getStatusCode(err1, restCountriesResp),
		CurrenciesApi:    getStatusCode(err2, currenciesResp),
	}
}

func getStatusCode(err error, res *http.Response) int {
	if err != nil {
		return http.StatusServiceUnavailable
	}
	return res.StatusCode
}
