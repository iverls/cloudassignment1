
# Country Info REST API

This project is an individual assignment that implements a REST API providing country information, historical population data, and the status of external API dependencies. It is written entirely in Go using standard libraries.

## Features

This REST API provides the following endpoints:

| Endpoint                                                       | Description                                                                                                                                            |
| -------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `/countryinfo/v1/info/{country}?limit={X}`                     | Retrieve general information about a country including name, population, languages, borders, flag URL, capital city, and cities (with optional limit). |
| `/countryinfo/v1/population/{country}?limit=startYear-endYear` | Retrieve historical population data for the given country with optional filtering by years.                                                            |
| `/countryinfo/v1/status/`                                      | Returns the status of external APIs and uptime of this service.                                                                                        |

## Technologies Used

- Go (Standard library)
- REST APIs (CountriesNow API, REST Countries API)

---

## Setup and Running Locally

### **Prerequisites:**

- [Go](https://golang.org/dl/) (Version 1.24 or newer)

### **Clone the Repository:**

```bash
git clone git@git.gvk.idi.ntnu.no:course/prog2005/prog2005-2025-workspace/iverls/assignment1_country_info.git
cd assignment1_country_info
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
GET http://localhost:8080/countryinfo/v1/info/no?limit=5
```

#### Example Response:

```json
{
  "name": "Norway",
  "continents": ["Europe"],
  "population": 5379475,
  "languages": {
    "nno": "Norwegian Nynorsk",
    "nob": "Norwegian Bokmål",
    "smi": "Sami"
  },
  "borders": ["FIN","SWE","RUS"],
  "flag": "https://flagcdn.com/w320/no.png",
  "capital": "Oslo",
  "cities": ["Abelvaer","Adalsbruk","Adland","Agotnes","Agskardet"]
}
```

### **Country Population:**

```http
GET http://localhost:8080/countryinfo/v1/population/no
```

#### Example Response:

```json
{
  "mean": 5238947,
  "values": [
    { "year": 2015, "value": 5188607 },
    { "year": 2016, "value": 5232929 },
    { "year": 2017, "value": 5265158 }
  ]
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

- The optional year-limit feature on the population endpoint is partially implemented; currently, all available data is returned.

---

## 🚀 Deployment

This service is deployed on [Render](https://render.com).

**Live URL:**\


## Repository



**Submitted by:**\
Iver Lieberg Stieng

