package main

//import "github.com/gin-gonic/gin"

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func initializeDataBase() (err error) {
	dsn := "root:yj950627@tcp(localhost:3309)/student?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect 2 database failed,err:%v", err)
		return nil
	}
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return nil
}

func main() {

	if err := initializeDataBase(); err != nil {
		panic(err)
	}
	fmt.Println("connect 2 database successs")
}
