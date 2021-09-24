package main

import (
	"database/sql"
	"fmt"
	_ "github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	_ "net/http"
)

type User struct{
	ID uint64 `json:"id"`
	Name string `json:"name"`
	Age uint32 `json:"age"`
}

var db *sql.DB

func initializeDatabase() (err error) {
	dsn := "root:yj950627@tcp(localhost:3309)/student?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return
	}
	if err = db.Ping(); err != nil {
		return
	}
	return nil
}

//查询单行数据
func querrySingleRowData() (user User){
	sqlStr := "select * from student where id = ?"
	if err := db.QueryRow(sqlStr,2).Scan(&user.ID,&user.Name,&user.Age);err!= nil{
		log.Printf("scan fail err : %v\n",err)
		return user
	}
	log.Println(user)
	return user
}

//查询多行数据
func queryMultiRowData() []User{
	sqlStr := "select * from student"
	rows,err := db.Query(sqlStr)
	if err != nil{
		log.Println(err)
		return nil
	}
	defer rows.Close()
	users := make([]User,0)
	for rows.Next(){
		var user User
		err := rows.Scan(&user.ID,&user.Name,&user.Age)
		if err != nil{
			log.Println(err)
			return nil
		}
		users = append(users,user)
	}
	return users
}

//更新数据
func updateRowData(){
	sqlStr := "update student set name = ? where id = ?"
	res,err := db.Exec(sqlStr,"王文皓",6)
	if err != nil{
		fmt.Sprintf("update faile err:%v\n",err)
	}
	//n:受影响的行数
	n,err := res.RowsAffected()
	if err != nil{
		fmt.Printf("get rows affexted failed")
		return
	}
	fmt.Printf("update success,affected %v\n",n)
}

func deleteRowData(){
	sqlStr := "delete from student where id = ?"
	res,err := db.Exec(sqlStr,3)
	if err != nil{
		fmt.Println(err)
		return
	}
	n,err := res.RowsAffected()
	if err != nil{
		fmt.Println("get rows affected failed")
	}
	fmt.Printf("delete success,affected:%v\n",n)
}

//插入数据
func InsertRowData(){
	sqlStr := "insert into student (name,age)values(?,?) "
	res,err := db.Exec(sqlStr,"wangbxu",24)
	if err != nil{
		fmt.Printf("insert failed err:%v\n",err)
		return
	}
	id,err := res.LastInsertId()
	if err != nil{
		fmt.Printf("get last insert id failed err :%v\n",err)
		return
	}
	fmt.Printf("get last insert id success id:%v",id)
}

func main() {
	if err := initializeDatabase(); err != nil {
		panic(err)
	}
	fmt.Println("connect 2 database success")
	//查询用户详情
	//r := gin.Default()
	////查询单行数据
	//r.GET("/user",func(c *gin.Context){
	//	user := querrySingleRowData()
	//	c.JSON(http.StatusOK,gin.H{
	//		"code": 0,
	//		"data":user,
	//	})
	//})
	////查询多行数据
	//r.GET("/users",func(c *gin.Context){
	//	user := queryMultiRowData()
	//	c.JSON(http.StatusOK,gin.H{
	//		"code": 0,
	//		"data":user,
	//	})
	//})
	//r.Run()

	//更新
	//updateRowData()
	//删除
	//deleteRowData()
	//插入
	InsertRowData()
}
