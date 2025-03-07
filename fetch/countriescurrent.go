package fetch

import (
	"github.com/iverls/assignment1_country_info/model"
)

func GetPopulationInfo(country string, limit string) (model.PopulationResponse, error) {
	req := map[string]string{"country": country}
	var data model.CountriesNowPopulation
	err := PostJSON("http://129.241.150.113:3500/api/v0.1/countries/population", req, &data)
	if err != nil {
		return model.PopulationResponse{}, err
	}
	return model.MapToPopulation(data, limit), nil
}

func GetCities(country string, limit int) ([]string, error) {
	req := map[string]string{"country": country}
	var resp struct {
		Error bool     `json:"error"`
		Msg   string   `json:"msg"`
		Data  []string `json:"data"`
	}

	err := PostJSON("http://129.241.150.113:3500/api/v0.1/countries/cities", req, &resp)
	if err != nil {
		return nil, err
	}

	if limit > 0 && len(resp.Data) > limit {
		return resp.Data[:limit], nil
	}

	return resp.Data, nil
}
