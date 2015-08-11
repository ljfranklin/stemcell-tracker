package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"strconv"
)

type StemcellHandler struct {
	stemcellMap map[string]map[string]string
}

func NewHandler() *StemcellHandler {
	return &StemcellHandler{
		stemcellMap: map[string]map[string]string{},
	}
}

func (h *StemcellHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	productName := r.URL.Query()["product_name"][0]
	productVersionParams := r.URL.Query()["product_version"]

	productVersion := "latest"
	if len(productVersionParams) > 0 {
		productVersion = productVersionParams[0]
	}

	switch r.Method {
	case "GET":
		versionMap := h.stemcellMap[productName]
		stemcell := versionMap[productVersion]
		fmt.Fprintf(w, stemcell)
	case "PUT":
		bodyContents, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if len(h.stemcellMap[productName]) == 0 {
			h.stemcellMap[productName] = map[string]string{}
		}

		if productVersion != "latest" {
			//check if this versioned stemcell is higher than latest
			latestVal, hasLatest := h.stemcellMap[productName]["latest"]
			if hasLatest == false {
				h.stemcellMap[productName]["latest"] = string(bodyContents)
			} else {
				latestStemcell, _ := strconv.Atoi(latestVal)
				versionStemcell, _ := strconv.Atoi(string(bodyContents))
				if versionStemcell > latestStemcell {
					h.stemcellMap[productName]["latest"] = string(bodyContents)
				}
			}
		}

		h.stemcellMap[productName][productVersion] = string(bodyContents)
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