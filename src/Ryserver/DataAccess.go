// DataAccess
package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func GetIdFromUserName(userName string) int {
	db, err := sql.Open("mysql", "ryan:1234@/rydb")

	var id int
	row := db.QueryRow("SELECT id FROM users WHERE name = ?", userName)
	err = row.Scan(&id)

	if err != nil {
		log.Fatal(err)
	}

	return id
}
