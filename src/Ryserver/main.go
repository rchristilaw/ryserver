// Ryserver project main.go
package main

import (
	"bytes"
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/index", Index)
	router.HandleFunc("/", Index)
	router.HandleFunc("/search/{searchVal}", SearchArtist)
	log.Fatal(http.ListenAndServe(":5555", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Index Page, %q", html.EscapeString(r.URL.Path))

}

func SearchArtist(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	searchVal := vars["searchVal"]
	fmt.Fprintln(w, "Search For:", searchVal)

	//url := "http://restapi3.apiary.io/notes"
	//fmt.Println("URL:>", url)

	var url = "https://itunes.apple.com/search?callback=callback"

	var searchEntity = "album"
	//var searchAttribute = "allArtistTerm"
	var searchLimit = "200"
	var searchTerm = searchVal

	url += "&entity=" + searchEntity
	//url += "&attribute=" + searchAttribute;
	url += "&limit=" + searchLimit
	url += "&term=" + searchTerm

	var jsonStr = []byte(`{"title":"Buy cheese and bread for breakfast."}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Fprintf(w, "response Status:", resp.Status)
	fmt.Fprintf(w, "response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Fprintf(w, "response Body:", string(body))

}
