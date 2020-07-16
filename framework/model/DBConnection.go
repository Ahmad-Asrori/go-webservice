package model

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)


var Db *sql.DB

func init()  {
	temp, err := sql.Open("mysql", "USER:PASSWORD@tcp(HOST:PORT)/DBNAME")
	if err != nil {
		log.Println(err)
	}

	err = temp.Ping()
	if err != nil {
		log.Println(err)
	}

	println("ping OK")

	Db = temp
}