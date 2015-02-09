// AlbumSearch
package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func SearchArtist(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	searchVal := vars["searchVal"]

	var url = "https://itunes.apple.com/search?"
	var searchEntity = "album"
	//var searchAttribute = "allArtistTerm"
	var searchLimit = "200"
	var searchTerm = searchVal

	url += "&entity=" + searchEntity
	//url += "&attribute=" + searchAttribute;
	url += "&limit=" + searchLimit
	url += "&term=" + searchTerm

	var jsonStr = []byte(`{"title":"Placeholder."}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	callback := r.FormValue("callback")

	body, _ := ioutil.ReadAll(resp.Body)

	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintf(w, "%s(%s)", callback, body)
}
