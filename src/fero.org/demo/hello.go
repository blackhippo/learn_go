package main

import (
	_ "github.com/go-sql-driver/mysql"

	"database/sql"
)

func main() {
	db, err := sql.Open("mysql", "root:admin@tcp(192.168.99.100:3306)/test")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("insert into account (name, password_hash, salt)"+
		"values (?, ?, ?)", "wei", "123", "234")
	if err != nil {
		panic(err)
	}
}
