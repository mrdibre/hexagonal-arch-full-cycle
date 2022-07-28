package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/mrdibre/hexagonal-arch-go/application"
)

type ProductDb struct {
	db *sql.DB
}

func (db *ProductDb) Get(id string) (application.ProductInterface, error) {
	var product application.Product

	stmt, err := db.db.Prepare("select id, name, price, status from products where id=?")

	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price, &product.Status)

	if err != nil {
		return nil, err
	}

	return &product, nil
}
