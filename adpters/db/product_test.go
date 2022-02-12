package db

import (
	"database/sql"
	"log"
)

var Db *sql.DB

func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
}

func createTable(db *sql.DB) {
	table := `CREATE TABLE products(
		"id": stirng,
		"name": string,
		"price": float,
		"status": string
		);`

	stmt, err := db.Prepare(table)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func createProduct(db *sql.DB) {
	insert := `insert into products valus("abc", "Product Test", 0, "DISABLED")`
	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}
