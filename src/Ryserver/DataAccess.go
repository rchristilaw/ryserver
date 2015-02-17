// DataAccess
package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func GetIdFromUserName(userName string) int {
	db, err := sql.Open("mysql", "ryan:1234@/rydb")

	var id int
	row := db.QueryRow("SELECT id FROM users WHERE name = ?", userName)
	err = row.Scan(&id)

	switch {
	case err == sql.ErrNoRows:
		log.Printf("No user with that ID.")
		return -1
	case err != nil:
		log.Fatal(err)
		return -1

	default:
		fmt.Printf("Username is %s\n", userName)
	}

	if err != nil {
		log.Fatal(err)
	}

	return id
}
