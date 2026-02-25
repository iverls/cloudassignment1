package main

import (
	"log"
	"net/http"
	"os"

	"github.com/iverls/assignment1_country_info/handler"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc(handler.ApiBase+handler.InfoPath, handler.CountryInfoHandler)
	http.HandleFunc(handler.ApiBase+handler.ExchangePath, handler.ExchangeHandler)
	http.HandleFunc(handler.ApiBase+handler.StatusPath, handler.StatusHandler)
	http.HandleFunc("/", handler.DefaultHandler)

	log.Println("Listening on port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
