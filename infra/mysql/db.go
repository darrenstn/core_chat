package mysql

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/core_chat?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
