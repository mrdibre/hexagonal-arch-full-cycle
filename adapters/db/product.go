package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/mrdibre/hexagonal-arch-go/application"
)

type ProductDb struct {
	db *sql.DB
}

func NewProductDb(db *sql.DB) *ProductDb {
	return &ProductDb{db}
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

func (db *ProductDb) create(product application.ProductInterface) (application.ProductInterface, error) {
	stmt, err := db.db.Prepare("INSERT INTO products(id, name, price, status) VALUES(?, ?, ?, ?);")

	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus())

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (db *ProductDb) update(product application.ProductInterface) (application.ProductInterface, error) {
	_, err := db.db.Exec(
		"UPDATE products SET name = ?, price = ?, status = ? WHERE id = ?",
		product.GetName(),
		product.GetPrice(),
		product.GetStatus(),
		product.GetID(),
	)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (db *ProductDb) Save(product application.ProductInterface) (application.ProductInterface, error) {
	var rows int

	db.db.QueryRow("SELECT id FROM products WHERE id = ?", product.GetID()).Scan(&rows)

	if rows == 0 {
		_, err := db.create(product)
		if err != nil {
			return nil, err
		}
	} else {
		_, err := db.update(product)
		if err != nil {
			return nil, err
		}
	}

	return product, nil
}
