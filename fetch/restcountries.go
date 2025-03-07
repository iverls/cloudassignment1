package fetch

import "github.com/iverls/assignment1_country_info/model"

func GetCountryInfo(code string) (model.CountryInfoResponse, error) {
	var data []model.RestCountriesResponse
	err := FetchJSON("http://129.241.150.113:8080/v3.1/alpha/"+code, &data)
	if err != nil || len(data) == 0 {
		return model.CountryInfoResponse{}, err
	}
	return model.MapToCountryInfo(data[0]), nil
}
