package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDatabase() *sql.DB {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "godb"
	// db, err = sql.Open("mysql", "<root>:<password>@/<dbname>")
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	// defer db.Close()

	return db
}
