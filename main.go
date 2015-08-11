package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

type StemcellHandler struct {
	stemcellMap map[string]string
}

func NewHandler() *StemcellHandler {
	return &StemcellHandler{
		stemcellMap: map[string]string{},
	}
}

func (h *StemcellHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	productName := r.URL.Query()["product_name"][0]

	switch r.Method {
	case "GET":
		stemcell := h.stemcellMap[productName]
		fmt.Fprintf(w, stemcell)
	case "PUT":
		bodyContents, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		h.stemcellMap[productName] = string(bodyContents)
		w.WriteHeader(http.StatusCreated)
	default:
		http.Error(w, "Method not found", http.StatusBadRequest)
	}
}

func main() {
	fmt.Println("Hello Gophers!")

	http.Handle("/stemcell", NewHandler())
	http.ListenAndServe(":8181", nil)
}
