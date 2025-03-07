package model

import (
	"strconv"
	"strings"
)

type RestCountriesResponse struct {
	Name struct {
		Common string `json:"common"`
	} `json:"name"`
	Continents []string          `json:"continents"`
	Population int               `json:"population"`
	Languages  map[string]string `json:"languages"`
	Borders    []string          `json:"borders"`
	Flags      struct {
		Png string `json:"png"`
	} `json:"flags"`
	Capital []string `json:"capital"`
}

type CountryInfoResponse struct {
	Name       string            `json:"name"`
	Continents []string          `json:"continents"`
	Population int               `json:"population"`
	Languages  map[string]string `json:"languages"`
	Borders    []string          `json:"borders"`
	Flag       string            `json:"flag"`
	Capital    string            `json:"capital"`
	Cities     []string          `json:"cities"`
}

func MapToCountryInfo(api RestCountriesResponse) CountryInfoResponse {
	capital := ""
	if len(api.Capital) > 0 {
		capital = api.Capital[0]
	}

	return CountryInfoResponse{
		Name:       api.Name.Common,
		Continents: api.Continents,
		Population: api.Population,
		Languages:  api.Languages,
		Borders:    api.Borders,
		Flag:       api.Flags.Png,
		Capital:    capital,
		Cities:     []string{},
	}
}

type CountriesNowPopulation struct {
	Error bool   `json:"error"`
	Msg   string `json:"msg"`
	Data  struct {
		Country          string `json:"country"`
		PopulationCounts []struct {
			Year  int `json:"year"`
			Value int `json:"value"`
		} `json:"populationCounts"`
	} `json:"data"`
}

type PopulationResponse struct {
	Mean   int                   `json:"mean"`
	Values []YearPopulationValue `json:"values"`
}

type YearPopulationValue struct {
	Year  int `json:"year"`
	Value int `json:"value"`
}

func MapToPopulation(api CountriesNowPopulation, limit string) PopulationResponse {
	var values []YearPopulationValue
	var sum int

	var startYear, endYear int
	if limit != "" {
		years := strings.Split(limit, "-")
		if len(years) == 2 {
			var err1, err2 error
			startYear, err1 = strconv.Atoi(years[0])
			endYear, err2 = strconv.Atoi(years[1])
			if err1 != nil || err2 != nil {
				startYear, endYear = 0, 0
			}
		}
	}

	for _, entry := range api.Data.PopulationCounts {
		if limit != "" && (entry.Year < startYear || entry.Year > endYear) {
			continue
		}
		values = append(values, YearPopulationValue{
			Year:  entry.Year,
			Value: entry.Value,
		})
		sum += entry.Value
	}

	var mean int
	if len(values) > 0 {
		mean = int(float64(sum)/float64(len(values)) + 0.5)
	}

	return PopulationResponse{
		Mean:   mean,
		Values: values,
	}
}

type StatusResponse struct {
	CountriesNowApi  int     `json:"countriesnowapi"`
	RestCountriesApi int     `json:"restcountriesapi"`
	Version          string  `json:"version"`
	Uptime           float64 `json:"uptime"`
}
