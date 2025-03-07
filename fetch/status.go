package fetch

import (
	"github.com/iverls/assignment1_country_info/model"
	"net/http"
)

func CheckApiStatus() model.StatusResponse {
	countriesNowResp, err1 := http.Get("http://129.241.150.113:3500/api/v0.1/")
	restCountriesResp, err2 := http.Get("http://129.241.150.113:8080/v3.1/alpha/no")

	return model.StatusResponse{
		CountriesNowApi:  getStatusCode(err1, countriesNowResp),
		RestCountriesApi: getStatusCode(err2, restCountriesResp),
	}
}

func getStatusCode(err error, res *http.Response) int {
	if err != nil {
		return http.StatusServiceUnavailable
	}
	return res.StatusCode
}
