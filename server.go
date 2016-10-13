package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	body, _ := ioutil.ReadAll(r.Body)
	switch r.Method {
	case "GET":
		resp = RespObj{RequestURI: r.RequestURI, Method: r.Method, Headers: r.Header, Host: r.Host}
	case "POST":
		var f interface{}
		_ = json.Unmarshal(body, &f)
		m := f.(map[string]interface{})
		resp = RespObj{RequestURI: r.RequestURI, Method: r.Method, Headers: r.Header, Host: r.Host, Params: m["params"].(string)}
	default:
	}
	fmt.Println(resp)
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
