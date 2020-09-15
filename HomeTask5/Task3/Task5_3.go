package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.Error(w, "404 page not found", http.StatusNotFound)
		return
	}

	if r.Method == http.MethodGet {
		http.ServeFile(w, r, "HomeTask5.html")
	}

	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		address := r.FormValue("address")
		token := fmt.Sprintf("%v:%v", name, address)
		cookie := http.Cookie{Name: "token", Value: token}
		http.SetCookie(w, &cookie)
		http.ServeFile(w, r, "HomeTask5.html")
	}
}

func main() {
	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(":8082", nil); err != nil {
		log.Fatal(err)
	}
}