package main

import (
	"encoding/json"
	"net/http"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
	case "POST":
	default:

	}
}

func main() {
	http.HandleFunc("/json", defaultHandler)
	http.ListenAndServe(":8000", nil)
}
