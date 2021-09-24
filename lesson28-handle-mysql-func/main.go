package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initializeDatabase() (err error) {
	dsn := "root:yj950627@tcp(localhost:3309)/mysql?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return
	}
	if err = db.Ping(); err != nil {
		return
	}
	return nil
}

func main() {
	if err := initializeDatabase(); err != nil {
		panic(err)
	}
	fmt.Println("connect 2 database success")
}
