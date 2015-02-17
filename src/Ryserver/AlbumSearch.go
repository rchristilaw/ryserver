// AlbumSearch
package main

import (
	"bytes"
	//	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type SearchParameters struct {
	Category string
	Value    string
}

func SearchArtist(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	searchVal := vars["searchVal"]
	log.Println("Search for " + searchVal)

	parameterMap := parseParameters(searchVal)

	url := prepareSearchUrl(parameterMap)

	log.Println("URL = " + url)

	var searchJsonData []byte

	if len(url) > 0 {
		searchJsonData = sendAppleSearchRequest(url)
	} else {
		searchJsonData = []byte(`{"resultsCount":0}`)
	}
	//searchJsonData := filterReleaseDates(searchJsonData)

	callback := r.FormValue("callback")
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s(%s)", callback, searchJsonData)
}

//func filterReleaseDates(jsonData []byte) []byte {

//	return jsonData
//}

func sendAppleSearchRequest(url string) []byte {
	log.Println(url)
	var jsonStr = []byte(`{"title":"Placeholder."}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	return body
}

func parseParameters(searchString string) []SearchParameters {

	log.Println("Parsing Parameters: " + searchString)

	params := strings.Split(searchString, "&")
	log.Println(len(params))
	parameterMap := make([]SearchParameters, len(params))

	for index, each := range params {
		keyValuePair := strings.Split(each, "=")
		log.Println(keyValuePair[0], keyValuePair[1])
		parameterMap[index] = SearchParameters{Category: keyValuePair[0], Value: keyValuePair[1]}
	}

	return parameterMap
}

func prepareSearchUrl(parameterMap []SearchParameters) string {
	var url = "https://itunes.apple.com/search?"
	var searchEntity = "album"
	//var searchAttribute = "allArtistTerm"
	var searchLimit = "200"

	url += "&entity=" + searchEntity
	//url += "&attribute=" + searchAttribute;
	url += "&limit=" + searchLimit

	for _, each := range parameterMap {
		if each.Category == "artist" {
			url += "&term=" + each.Value
		}
	}

	return url

}
