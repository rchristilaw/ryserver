// Ryserver project main.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func main() {
	InitRoutes()
}

type UserInfo struct {
	UserName string
	Password string
}

func Login(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	loginParam := vars["loginParam"]
	userInfo := parseUserInformation(loginParam)
	callback := r.FormValue("callback")
	id := GetIdFromUserName(userInfo.UserName)

	if id == -1 {
		fmt.Fprintf(w, "Login Failed", callback, id)
	} else {
		fmt.Fprintf(w, "Login Successful", callback, id)
	}
}

func parseUserInformation(userInfoString string) UserInfo {

	log.Println("Login Info: " + userInfoString)

	loginInfo := ParseHttpParameters(userInfoString)

	keyValuePair := ParseKeyValue(loginInfo[0])

	name := keyValuePair[1]

	keyValuePair = ParseKeyValue(loginInfo[1])

	password := keyValuePair[1]

	return UserInfo{name, password}
}

func ParseHttpParameters(httpParams string) []string {
	return strings.Split(httpParams, "&")
}

func ParseKeyValue(parameter string) []string {
	return strings.Split(parameter, "=")
}
