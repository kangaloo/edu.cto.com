package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

var (
	db  *sqlx.DB
	err error
)

func init() {
	db, err = sqlx.Open("mysql",
		"go:123456@tcp(localhost:3306)/test")
	if err != nil {
		db.Close()
		log.Fatalln(err)
	}
}

func main() {

	//rows, err := db.Query("select user_id, username, sex, email from person")

	if err != nil {
		log.Fatalln(err)
	}

	var p []*Person
	db.Select(&p, "select  user_id, username, sex, email from person")

	for _, v := range p {
		fmt.Println(v.Email)
	}

}

type Person struct {
	Id    int    `db:"user_id"`
	Name  string `db:"username"`
	Sex   string `db:"sex"`
	Email string `db:"email"`
}
