package main

import (
	"fmt"
	"log"
)

// QueryRow会建立连接，并在调用scan时释放，可以用SetMaxOpenConns和SetMaxIdleConns调整
func getOne(id int) (app, error) { //查
	a := app{}
	row := db.QueryRow("select id, name, age from user where id=?", id) //mysql的写法是?配合参数，该函数后面最好直接跟scan以此释放连接
	err := row.Scan(&a.id, &a.name, &a.age)

	return a, err
}

func getMany(id int) (apps []app, err error) { //查
	rows, err := db.Query("select id, name, age from user where id>=?", id)
	defer rows.Close() //rows要close，而row不用，因为row在scan后就关闭，而rows在scan后不会关闭
	if err != nil {
		log.Fatalln(err.Error())
	}

	for rows.Next() {
		a := app{}
		err = rows.Scan(&a.id, &a.name, &a.age)
		apps = append(apps, a)
	}

	return apps, err
}

// 插入，更新和删除全部用Exec方法 func (db *DB) Exec(query string, args ...interface{}) (Result, error)
func inserRow() {
	res, err := db.Exec("insert into user(name, age) values(?, ?)", "ho", "23")
	if err != nil {
		log.Fatalln(err.Error())
	}

	theID, err := res.LastInsertId()
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", theID)
}

func updateRow() {
	res, err := db.Exec("update user set name='hobbyGG', age=24 where id = 2")
	if err != nil {
		log.Fatalln(err.Error())
	}

	n, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("get RowsAffected failed, err:%v\n", err)
	}
	fmt.Printf("update success, affected rows:%d\n", n)
}

func deleteRow() {
	res ,err := db.Exec("delete from user where name='ho'")
	if err != nil {
		log.Fatalf("get RowsAffected failed, err:%v\n", err)
	}
	
	n, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("get RowsAffected failed, err:%v\n", err)
	}
	fmt.Printf("update success, affected rows:%d\n", n)
}