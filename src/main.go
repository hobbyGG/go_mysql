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
	database  = "sql_test"
)

func main() {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user_name, password, ip, port, database) //连接字符串

	var err error

	//不可以用:=，因为是给全局变量赋值
	db, err = sql.Open("mysql", connStr) //数据库驱动，连接字符串。注意，open只检查参数合法性，并不建立连接
	if err != nil {
		log.Fatalln(err.Error()) //Fatalln就是println再exit
	}
	defer db.Close() //这行要在if后面，因为如果在if前面，则会在panic退出时关闭db，而此时db是空的，会出错

	ctx := context.Background() /*需要一个上下文类型的参数，context可以携带请求范围的值
	该方法是返回一个非nil的空context*/
	err = db.PingContext(ctx) //验证连接是否有效
	if err != nil {
		log.Fatalln(err.Error())
	}
	//以上直接写成err := db.Ping()也可以

	fmt.Println("Connected!")

	inserRow()
	updateRow()
	deleteRow()

	one, err := getOne(0)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(one)

	many, err := getMany(1)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(many)
}
