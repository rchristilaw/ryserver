// RoutingManager
package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func InitRoutes() {
	log.Println("Initializing routes")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/sendemail", SendEmail)
	router.HandleFunc("/login/{loginParam}", Login)
	router.HandleFunc("/search/{searchVal}", SearchArtist)
	log.Fatal(http.ListenAndServe(":5555", router))
}
