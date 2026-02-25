package model

// --- REST Countries API Internal Models ---

type RestCountriesResponse struct {
	Name struct {
		Common string `json:"common"`
	} `json:"name"`
	Continents []string          `json:"continents"`
	Population int               `json:"population"`
	Area       float64           `json:"area"`
	Languages  map[string]string `json:"languages"`
	Borders    []string          `json:"borders"`
	Flags      struct {
		Png string `json:"png"`
	} `json:"flags"`
	Capital    []string `json:"capital"`
	Currencies map[string]struct {
		Name   string `json:"name"`
		Symbol string `json:"symbol"`
	} `json:"currencies"`
}

// --- Currency API Internal Models ---

type CurrencyApiResponse struct {
	Base  string             `json:"base"`
	Rates map[string]float64 `json:"rates"`
}

// --- API Output Models (What your API returns) ---

type CountryInfoResponse struct {
	Name       string            `json:"name"`
	Continents []string          `json:"continents"`
	Population int               `json:"population"`
	Area       float64           `json:"area"`
	Languages  map[string]string `json:"languages"`
	Borders    []string          `json:"borders"`
	Flag       string            `json:"flag"`
	Capital    string            `json:"capital"`
}

type ExchangeResponse struct {
	Country       string             `json:"country"`
	BaseCurrency  string             `json:"base-currency"`
	ExchangeRates map[string]float64 `json:"exchange-rates"`
}

type StatusResponse struct {
	RestCountriesApi int     `json:"restcountriesapi"`
	CurrenciesApi    int     `json:"currenciesapi"`
	Version          string  `json:"version"`
	Uptime           float64 `json:"uptime"`
}
