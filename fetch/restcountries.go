package fetch

import "github.com/iverls/assignment1_country_info/model"

// GetCountryData fetches raw data for a country code.
func GetCountryData(code string) (model.RestCountriesResponse, error) {
	var data []model.RestCountriesResponse
	err := FetchJSON("http://129.241.150.113:8080/v3.1/alpha/"+code, &data)
	if err != nil || len(data) == 0 {
		return model.RestCountriesResponse{}, err
	}
	return data[0], nil
}
