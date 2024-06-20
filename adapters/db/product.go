package db

import (
	"database/sql"

	"github.com/DerivedPuma7/go-hexagonal/application"
	"github.com/DerivedPuma7/go-hexagonal/application/interfaces"
	_ "github.com/mattn/go-sqlite3"
)

type ProductDb struct {
	db *sql.DB
}

func NewProductDb(db *sql.DB) *ProductDb {
	return &ProductDb{db:db}
}

var _ interfaces.ProductPersistenceInterface = (*ProductDb)(nil)

func (p *ProductDb) Get(id string) (interfaces.ProductInterface, error) {
	var product application.Product
	stmt, err := p.db.Prepare("select id, name, price, status from products where id=?")
	if err != nil {
		return nil, err
	}
	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price, &product.Status)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (p *ProductDb) Save(product interfaces.ProductInterface) (interfaces.ProductInterface, error) {
	exists, err := p.productExists(product.GetID())
	if err != nil {
		return nil, err
	}

	if exists {
		_, err := p.update(product)
		if err != nil {
			return nil, err
		}
	} else {
		_, err := p.create(product)
		if err != nil {
			return nil, err
		}
	}

	return product, nil
}

func (p *ProductDb) productExists(id string) (bool, error) {
	var existingID string
	err := p.db.QueryRow("select id from products where id=?", id).Scan(&existingID)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (p *ProductDb) create(product interfaces.ProductInterface) (interfaces.ProductInterface, error) {
	stmt, err := p.db.Prepare(`insert into products(id, name, price, status) values(?, ?, ?, ?)`)
	if err != nil {
		return nil, err
	}
	_, err = stmt.Exec(
		product.GetID(),
		product.GetName(),
		product.GetPrice(),
		product.GetStatus(),
	)
	if err != nil {
		return nil, err
	}
	err = stmt.Close()
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *ProductDb) update(product interfaces.ProductInterface) (interfaces.ProductInterface, error) {
	_, err := p.db.Exec(`update products set name=?, price=?, status=? where id=?`, product.GetName(), product.GetPrice(), product.GetStatus(), product.GetID())
	if err != nil {
		return nil, err
	}
	return product, nil
}

