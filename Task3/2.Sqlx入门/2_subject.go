package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type book struct {
	Id     int     `db:"id"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float32 `db:"price"`
}

func test2Main(db *sqlx.DB) {
	books := []book{}
	db.Select(&books, "SELECT * FROM books WHERE price > ?", 50)
	fmt.Println(books)
}
