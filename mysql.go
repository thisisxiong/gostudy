package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Person struct {
	UserId   int    `db:"user_id"`
	Username string `db:"username"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}

type Place struct {
	Country string `db:"country"`
	City    string `db:"city"`
	TelCode int    `db:"telcode"`
}

var Db *sqlx.DB

func init() {
	database, err := sqlx.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println("##open mysql failed:", err)
		return
	}
	Db = database
}

func main() {
	//insert()
	//fetch(1)
	//update()
	del(4)
}

func insert() {
	r, err := Db.Exec("insert into person(username,sex,email) values(?,?,?)", "study1", "man", "admin@11.com")
	if err != nil {
		fmt.Println("##exec failed:", err)
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("###exec failed:", err)
		return
	}
	defer Db.Close()
	fmt.Println("insert succ:", id)
}

func fetch(id int) {
	var person []Person
	err := Db.Select(&person, "select user_id,username,sex,email from person where user_id=?", id)
	if err != nil {
		fmt.Println("##Select failed:", err)
		return
	}
	fmt.Println("select succ:", person)
}

func update() {
	res, err := Db.Exec("update person set username =? where user_id = ? ", "jack", 1)
	if err != nil {
		fmt.Println("update failed:", err)
		return
	}
	row, err := res.RowsAffected()
	if err != nil {
		fmt.Println("rows failed:", err)
	}
	fmt.Println("update succ:", row)
}

func del(id int) {
	res, err := Db.Exec("delete from person where user_id=? ", id)
	if err != nil {
		fmt.Println("delete failed:", err)
		return
	}
	row, err := res.RowsAffected()
	if err != nil {
		fmt.Println("row failed:", err)
	}
	fmt.Println("delete succ:", row)
}
