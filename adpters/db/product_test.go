package db_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/matheusr42/go-hexagonal/adpters/db"
	"github.com/matheusr42/go-hexagonal/application"
	"github.com/stretchr/testify/require"
)

var Db *sql.DB

func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	createProduct(Db)
}

func createTable(db *sql.DB) {
	table := `CREATE TABLE products(
		id stirng,
		name string,
		price float,
		status int
		);`

	stmt, err := db.Prepare(table)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func createProduct(db *sql.DB) {
	insert := `insert into products values("abc", "Product Test", 0, 0)`
	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func TestProductDb_Get(t *testing.T) {
	setUp()
	defer Db.Close()

	productDb := db.NewProductDb(Db)

	product, err := productDb.Get("abc")

	require.Nil(t, err)
	require.Equal(t, "Product Test", product.GetName())
	require.Equal(t, 0.0, product.GetPrice())
	require.Equal(t, application.DISABLED, product.GetStatus())
}

func TestProductDb_Save(t *testing.T) {
	setUp()
	defer Db.Close()
	productDb := db.NewProductDb(Db)

	product := application.NewProduct()
	product.Name = "Test 1"
	product.Price = 10

	result, err := productDb.Save(product)
	require.Nil(t, err)
	require.Equal(t, product.Name, result.GetName())
	require.Equal(t, product.Price, result.GetPrice())
	require.Equal(t, *product.Status, result.GetStatus())

	status := application.ENABLED
	product.Status = &status
	result, err = productDb.Save(product)
	require.Nil(t, err)
	require.Equal(t, product.Name, result.GetName())
	require.Equal(t, product.Price, result.GetPrice())
	require.Equal(t, *product.Status, result.GetStatus())
}
