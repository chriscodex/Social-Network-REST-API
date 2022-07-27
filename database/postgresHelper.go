package database

import (
	"database/sql"
	"log"
)

func CloseReadingRows(r *sql.Rows) {
	err := r.Close()
	if err != nil {
		log.Fatal(err)
	}
}
