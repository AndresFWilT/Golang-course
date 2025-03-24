package storage

import (
	"database/sql"
	"github.com/AndresFWilT/postgres-db/package/invoiceheader"
)

const (
	psqlCreateInvoiceHeader = `INSERT INTO invoice_headers(client) VALUES($1) RETURNING id, created_at`
)

// PsqlInvoiceHeader used for work with postgres - invoiceHeader
type PsqlInvoiceHeader struct {
	db *sql.DB
}

// NewPsqlInvoiceHeader return a new pointer of PsqlInvoiceHeader
func NewPsqlInvoiceHeader(db *sql.DB) *PsqlInvoiceHeader {
	return &PsqlInvoiceHeader{db}
}

func (p *PsqlInvoiceHeader) CreateTx(tx *sql.Tx, m *invoiceheader.Model) error {
	stmt, err := tx.Prepare(psqlCreateInvoiceHeader)
	if err != nil {
		return err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {

		}
	}(stmt)

	return stmt.QueryRow(m.Client).Scan(&m.ID, &m.CreatedAt)
}
