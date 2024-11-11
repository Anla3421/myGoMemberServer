package dao

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var MysqlConn *sql.DB

func init() {
	fmt.Println("MySQL initial")
	CreateConn()
}

func CreateConn() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/zmemberdb?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	MysqlConn = db
	fmt.Println("Suceessfully Connected to MySQL")
}
