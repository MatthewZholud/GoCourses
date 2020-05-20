package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type headers struct {
	Accept     []string `json:"Accept"`
	User_Agent []string `json:"User-Agent"`
}

type request struct {
	Host        string  `json:"Host"`
	User_agent  string  `json:"User_agent"`
	Request_uri string  `json:"Request_uri"`
	Headers     headers `json:"Headers"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	var rq request
	rq.Host = r.Host
	rq.User_agent = r.UserAgent()
	rq.Request_uri = r.RequestURI
	rq.Headers.Accept = r.Header["Accept"]
	rq.Headers.User_Agent = r.Header["User_Arent"]

	RqJson, err := json.Marshal(rq)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(RqJson)
}

func main() {
	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}