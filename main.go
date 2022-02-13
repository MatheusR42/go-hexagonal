package main

import (
	"database/sql"

	"github.com/matheusr42/go-hexagonal/adpters/db"
	"github.com/matheusr42/go-hexagonal/application"
)

func main() {
	dbLocal, _ := sql.Open("sqlite3", "db.sqlite")
	productDbAdpter := db.NewProductDb(dbLocal)
	productService := application.NewProductService(productDbAdpter)

	product, _ := productService.Create("Product 1", 30)
	productService.Enable(product)
}
