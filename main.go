package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	productName := r.URL.Query()["product_name"][0]

	switch r.Method {
	case "GET":
		stemcell := "3030"
		if productName == "cf-mysql" {
			stemcell = "3026"
		}

		fmt.Fprintf(w, stemcell)
	case "PUT":
		w.WriteHeader(http.StatusCreated)
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func main() {
	fmt.Println("Hello Gophers!")

	http.HandleFunc("/stemcell", handler)
	http.ListenAndServe(":8181", nil)
}
