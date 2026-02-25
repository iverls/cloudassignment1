
# Country Info REST API

This project is an individual assignment that implements a REST API providing country information, currency exchange rates for neighboring countries, and the status of external API dependencies. It is written entirely in Go using standard libraries.

## Features

This REST API provides the following endpoints:

| Endpoint                                    | Description                                                                                                                                              |
| ------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `/countryinfo/v1/info/{country}`            | Retrieve general information about a country including name, continents, population, area, languages, borders, flag URL, and capital city.               |
| `/countryinfo/v1/exchange/{country}`        | Retrieve currency exchange rates from the requested country's base currency to the currencies used by all of its bordering neighboring countries.        |
| `/countryinfo/v1/status/`                   | Returns the HTTP status codes of the external APIs and the total uptime of this service.                                                                 |

## Technologies Used

- Go (Standard library only - no third-party dependencies)
- External REST APIs:
  - **REST Countries API** (`http://129.241.150.113:8080/v3.1/`)
  - **Currency API** (`http://129.241.150.113:9090/currency/`)

---

## Setup and Running Locally

### **Prerequisites:**

- [Go](https://golang.org/dl/) (Version 1.22 or newer)

### **Clone the Repository:**

```bash

```

### **Run the Service:**

```bash
go mod tidy
go build ./...
go run ./terminal/main.go
```

The server will start listening on port `8080`.

---

## Endpoint Examples

### **Country Information:**

```http
GET http://localhost:8080/countryinfo/v1/info/no
```

#### Example Response:

```json
{
  "name": "Norway",
  "continents": [
    "Europe"
  ],
  "population": 5379475,
  "area": 323802,
  "languages": {
    "nno": "Norwegian Nynorsk",
    "nob": "Norwegian Bokmål",
    "smi": "Sami"
  },
  "borders": [
    "FIN",
    "SWE",
    "RUS"
  ],
  "flag": "[https://flagcdn.com/w320/no.png](https://flagcdn.com/w320/no.png)",
  "capital": "Oslo"
}
```

### **Neighboring Exchange Rates:**

```http
GET http://localhost:8080/countryinfo/v1/exchange/no
```

#### Example Response:

```json
{
  "country": "Norway",
  "base-currency": "NOK",
  "exchange-rates": {
    "EUR": 0.08878,
    "RUB": 8.011422,
    "SEK": 0.946335
  }
}
```

### **API Status:**

```http
GET http://localhost:8080/countryinfo/v1/status/
```

#### Example Response:

```json
{
  "countriesnowapi": 200,
  "restcountriesapi": 200,
  "version": "v1",
  "uptime": 142.532
}
```

---

## Known Issues



## 🚀 Deployment

This service is deployed on [Render](https://render.com).

**Live URL:**\


## Repository



**Submitted by:**\
Iver Lieberg Stieng

