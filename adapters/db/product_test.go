package db_test

import (
	"database/sql"
	"github.com/mrdibre/hexagonal-arch-go/adapters/db"
	"github.com/mrdibre/hexagonal-arch-go/application"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

var Db *sql.DB

func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	createProduct(Db)
}

func createTable(db *sql.DB) {
	table := "CREATE TABLE PRODUCTS(id string, name string, price float, status string);"

	stmt, err := Db.Prepare(table)

	if err != nil {
		log.Fatalln(err.Error())
	}

	stmt.Exec()
}

func createProduct(db *sql.DB) {
	insert := `INSERT INTO products VALUES("abc", "Product Test", 0, "disabled");`

	stmt, err := db.Prepare(insert)

	if err != nil {
		log.Fatalln(err.Error())
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
	product.Name = "Product Test"
	product.Price = 25

	result, err := productDb.Save(product)

	require.Nil(t, err)
	require.Equal(t, product.Name, result.GetName())
	require.Equal(t, product.Price, result.GetPrice())
	require.Equal(t, product.Status, result.GetStatus())

	product.Enable()

	result, err = productDb.Save(product)

	require.Nil(t, err)
	require.Equal(t, product.Status, result.GetStatus())
}
