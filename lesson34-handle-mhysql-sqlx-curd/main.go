package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type User struct {
	ID   uint64 `db:"id"`
	Name string `db:"name"`
	Age  uint32 `db:"age"`
}

var db *sqlx.DB

func initializeDataBase() (err error) {
	dsn := "root:yj950627@tcp(localhost:3309)/student?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect 2 database failed,err:%v", err)
		return
	}
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return nil
}

func querySingleRowData() {
	sqlStr := "select * from student where id = ?"
	var u User
	if err := db.Get(&u, sqlStr, 7); err != nil {
		fmt.Printf("err : %v", err)
		return
	}
	fmt.Printf("user :%v", u)
}

func queryMultiRowData() []User {
	sqlStr := "select * from student"
	var users []User
	if err := db.Select(&users, sqlStr); err != nil {
		fmt.Printf("query failed,err:%v", err)
		return nil
	}
	return users
}

func updateRowData() {
	sqlStr := "update student set age = ? where id = ?"
	res, err := db.Exec(sqlStr, 33, 3)
	if err != nil {
		fmt.Printf("update fail ,err :%v\n", err)
		return
	}
	n, err := res.RowsAffected()
	if err != nil {
		fmt.Printf("get RowsAffected failed,err:%v", err)
		return
	}
	fmt.Printf("effected num:%v", n)

}

func insertRowData() {
	sqlStr := "insert into student (name,age) values(?,?)"
	res, err := db.Exec(sqlStr, "进字节", 25000)
	if err != nil {
		fmt.Printf("insert failed err:%v", err)
		return
	}
	n, err := res.LastInsertId()
	if err != nil {
		fmt.Printf("get LastInsertId failed err:%v", err)
		return
	}
	fmt.Printf("insert success,last insert id:%v", n)
}

func deleteRowData() {
	sqlStr := "delete from student where id = ?"
	res, err := db.Exec(sqlStr, 10)
	if err != nil {
		fmt.Printf("delete failed err:%v", err)
		return
	}
	n, err := res.RowsAffected()
	if err != nil {
		fmt.Printf("get rows affected failed,err:%v", err)
		return
	}
	fmt.Printf("effected num %v", n)
}

func selectNamedQuery() {
	sqlStr := "select * from student where age = :age"
	//rows, err := db.NamedQuery(sqlStr, map[string]interface{}{
	//	"age": 25000,
	//})
	user := User{
		Age: 31,
	}
	rows, err := db.NamedQuery(sqlStr, user)
	if err != nil {
		fmt.Printf("named query failed,err:%v", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var u User
		if err := rows.StructScan(&u); err != nil {
			fmt.Printf("struct sacn failed err :%v", err)
			//return
			continue
		}
		fmt.Println(u)
	}
}

func batchInsert() {
	users := []User{
		{Name: "111", Age: 100},
		{Name: "222", Age: 100},
		{Name: "333", Age: 100},
	}
	sqlStr := "insert into student (name,age) values (:name,:age)"
	_, err := db.NamedExec(sqlStr, users)
	if err != nil {
		fmt.Printf("batch insert failed,err:%v\n", err)
		return
	}
	fmt.Printf("batch insert success\n")

}

func main() {
	if err := initializeDataBase(); err != nil {
		panic(err)
	}
	//querySingleRowData()
	//fmt.Println("connect 2 database success")

	//r := gin.Default()
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSONP(http.StatusOK, gin.H{
	//		"message": "pong",
	//		"data":    queryMultiRowData(),
	//	})
	//})
	//r.Run()

	//updateRowData()

	//insertRowData()

	//deleteRowData()

	//selectNamedQuery()

	batchInsert()
}
