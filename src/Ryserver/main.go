// Ryserver project main.go
package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	InitRoutes()
}

func Login(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userName := vars["userName"]
	id := GetIdFromUserName(userName)
	fmt.Fprintf(w, "Login Successful", id)
}
