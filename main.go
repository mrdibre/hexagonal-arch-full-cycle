package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/mrdibre/hexagonal-arch-go/adapters/db"
	"github.com/mrdibre/hexagonal-arch-go/application"
)

func main() {
	connection, _ := sql.Open("sqlite3", "db.sqlite")
	productPersistenceAdapter := db.NewProductDb(connection)
	productService := application.NewProductService(productPersistenceAdapter)

	product, _ := productService.Create("Product 2", 100)
	productService.Enable(product)
}
