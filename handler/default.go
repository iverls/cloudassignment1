package handler

import "net/http"

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Invalid endpoint. Please refer to API documentation.", http.StatusNotFound)
}
