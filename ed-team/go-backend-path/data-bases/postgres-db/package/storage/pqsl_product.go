package storage

import (
	"database/sql"
	"fmt"
	"github.com/AndresFWilT/postgres-db/package/product"
)

const (
	psqlCreateProduct = `INSERT INTO products(name, observations, price, created_At) VALUES($1, $2, $3, $4) RETURNING id`
	psqlGetAllProduct = `SELECT id, name, observations, price, 
	created_at, updated_at
	FROM products`
	psqlGetProductByID = psqlGetAllProduct + " WHERE id = $1"
	psqlUpdateProduct  = `UPDATE products SET name = $1, observations = $2,
	price = $3, updated_at = $4 WHERE id = $5`
	psqlDeleteProduct = `DELETE FROM products WHERE id = $1`
)

type PsqlProduct struct {
	db *sql.DB
}

func NewPsqlProduct(db *sql.DB) *PsqlProduct {
	return &PsqlProduct{db}
}

func (p *PsqlProduct) Create(m *product.Model) error {
	stmt, err := p.db.Prepare(psqlCreateProduct)
	if err != nil {
		return err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {

		}
	}(stmt)

	err = stmt.QueryRow(
		m.Name,
		stringToNull(m.Observations),
		m.Price,
		m.CreatedAt,
	).Scan(&m.ID)
	if err != nil {
		return err
	}

	fmt.Println("Product created correctly")
	return nil
}

func (p *PsqlProduct) GetAll() (product.Models, error) {
	stmt, err := p.db.Prepare(psqlGetAllProduct)
	if err != nil {
		return nil, err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {

		}
	}(stmt)

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	ms := make(product.Models, 0)
	for rows.Next() {
		m := &product.Model{}
		observationNull := sql.NullString{}
		updatedAtNull := sql.NullTime{}

		err := rows.Scan(
			&m.ID,
			&m.Name,
			&observationNull,
			&m.Price,
			&m.CreatedAt,
			&updatedAtNull,
		)
		if err != nil {
			return nil, err
		}

		m.Observations = observationNull.String
		m.UpdatedAt = updatedAtNull.Time
		ms = append(ms, *m)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return ms, nil
}

func (p *PsqlProduct) GetById(id uint) (*product.Model, error) {
	stmt, err := p.db.Prepare(psqlGetProductByID)
	if err != nil {
		return &product.Model{}, err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {

		}
	}(stmt)

	return scanRowProduct(stmt.QueryRow(id)) // helper in the storage
}

func (p *PsqlProduct) Update(m *product.Model) error {
	stmt, err := p.db.Prepare(psqlUpdateProduct)
	if err != nil {
		return err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {

		}
	}(stmt)

	res, err := stmt.Exec(
		m.Name,
		stringToNull(m.Observations), // null control
		m.Price,
		timeToNull(m.UpdatedAt),
		m.ID,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("product with id: %d does not exist", m.ID)
	}

	fmt.Println("Product updated correctly")
	return nil
}

func (p *PsqlProduct) Delete(id uint) error {
	stmt, err := p.db.Prepare(psqlDeleteProduct)
	if err != nil {
		return err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {

		}
	}(stmt)

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	fmt.Println("Product deleted correctly")
	return nil
}
