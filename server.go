package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type RespObj struct {
	RequestURI string
	Method     string
	Headers    http.Header
	Host       string
	Params     string
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	var resp RespObj
	switch r.Method {
	case "GET":
		resp = RespObj{RequestURI: r.RequestURI, Method: r.Method, Headers: r.Header, Host: r.Host}
	case "POST":
		r.ParseForm()
		resp = RespObj{RequestURI: r.RequestURI, Method: r.Method, Headers: r.Header, Host: r.Host}
		resp.Params = r.FormValue("params")
	default:
	}
	sendResponse(resp, w)
}

func sendResponse(resp RespObj, w http.ResponseWriter) {
	js, err := json.Marshal(resp)
	if err == nil {
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	} else {
		fmt.Println(err)
	}
}

func main() {
	http.HandleFunc("/json", defaultHandler)
	http.ListenAndServe(":8000", nil)
}
