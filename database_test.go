package belajar_golang_database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestEmpty(t *testing.T) {

}

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/golang_database")
	if err != nil {
		panic(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
		if err != nil {

		}
	}(db)
}
