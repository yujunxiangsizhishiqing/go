package main

import (
	"database/sql"
	"fmt"

	//_ "github.com/go-sql-driver/mysql"
	_ "gorm.io/driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:yj950627@tcp(localhost:3309)/mysql")
	//fmt.Println(db)
	//fmt.Println(err)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	//尝试建立链接
	if err := db.Ping(); err != nil {
		fmt.Println("connect 2 database failed")
		panic(err)
	}
	//设置最大链接数
	db.SetMaxOpenConns(10)
	//设置最大闲置链接数
	db.SetMaxIdleConns(10)
	//执行到这里说明连接成功
	fmt.Println("connect 2 database success")
}
