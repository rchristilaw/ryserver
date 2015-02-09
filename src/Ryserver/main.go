// Ryserver project main.go
package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/index", Index)
	router.HandleFunc("/login/{userName}", Login)
	router.HandleFunc("/search/{searchVal}", SearchArtist)
	log.Fatal(http.ListenAndServe(":5555", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Index Page, %q", html.EscapeString(r.URL.Path))

}

func Login(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userName := vars["userName"]

	db, err := sql.Open("mysql", "ryan:1234Abcd!@/rydb")

	var id int64
	row := db.QueryRow("SELECT id FROM users WHERE name = ?", userName)
	err = row.Scan(&id)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "Login Successful", id)
}

func SearchArtist(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	searchVal := vars["searchVal"]

	//var url = "https://itunes.apple.com/search?callback=callback"
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
