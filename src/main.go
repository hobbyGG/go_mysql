package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" //这个包有个init会自我注册，下划线的意思是给这个包一个名字
)

var db *sql.DB //数据库

const (
	user_name = "root"
	password  = "1123"
	ip        = "127.0.0.1"
	port      = "6666"
	database  = "godb"
)

func main() {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user_name, password, ip, port, database) //连接字符串

	db, err := sql.Open("mysql", connStr) //数据库驱动，连接字符串
	if err != nil {
		log.Fatalln(err.Error())	//Fatalln就是println再exit
	}

	ctx := context.Background() /*需要一个上下文类型的参数，context可以携带请求范围的值
	该方法是返回一个非nil的空context*/

	err = db.PingContext(ctx) //验证连接是否有效
	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Println("Connected!")
}
