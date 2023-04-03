package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const createBooksTable = `
CREATE TABLE IF NOT EXISTS books (
	id SERIAL PRIMARY KEY,
	title TEXT,
	author TEXT,
    publisher TEXT,
    publish_date DATE NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
)
`

func migration() {
	db, err := sql.Open("postgres", "postgres://user:password@localhost/bookstore?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec(createBooksTable)
	if err != nil {
		panic(err)
	}

	fmt.Println("Migration completed successfully")
}
