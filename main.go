package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"strconv"
	"os"
)

type StemcellHandler struct {
	stemcellMap map[string]map[string]string
}

func NewStemcellHandler(stemcellMap map[string]map[string]string) *StemcellHandler {
	return &StemcellHandler{
		stemcellMap: stemcellMap,
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

type BadgeHandler struct {
	stemcellMap map[string]map[string]string
}

func NewBadgeHandler(stemcellMap map[string]map[string]string) *BadgeHandler {
	return &BadgeHandler{
		stemcellMap: stemcellMap,
	}
}

func (h *BadgeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	productName := r.URL.Query()["product_name"][0]
	productVersion := "latest"

	switch r.Method {
	case "GET":
		versionMap, foundProduct := h.stemcellMap[productName]
		if foundProduct == false {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		stemcell := versionMap[productVersion]
		badgeUrl := fmt.Sprintf("https://img.shields.io/badge/stemcell-%s-brightgreen.svg", stemcell)

		fmt.Printf("Started badge request: %s\n", badgeUrl)
		badgeResp, err := http.Get(badgeUrl)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error fetching badge: %s", err), http.StatusInternalServerError)
			return
		}
		fmt.Printf("Received badge response: %d\n", badgeResp.StatusCode)

		badgeContents, err := ioutil.ReadAll(badgeResp.Body)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error fetching badge: %s", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "image/svg+xml")
		w.Write(badgeContents)
	default:
		http.Error(w, "Method not found", http.StatusBadRequest)
	}
}

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8181"
	}

	stemcellMap := map[string]map[string]string{}
	http.Handle("/stemcell", NewStemcellHandler(stemcellMap))
	http.Handle("/badge", NewBadgeHandler(stemcellMap))

	fmt.Printf("Stemcell tracker is listening on port %s...\n", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}