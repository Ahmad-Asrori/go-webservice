package Model

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
)

var Db *sql.DB

func init() {
	tempDb, err := sql.Open("mysql", "USERNAME:PASSWORD@tcp(HOST:POST)/DBNAME")
	if err != nil {
		fmt.Println(err.Error())
	}

	Db = tempDb
}

