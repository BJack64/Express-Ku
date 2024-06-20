package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Init() *sql.DB {
	dsn := "root@tcp(localhost:3306)/express-ku"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	return db
}
