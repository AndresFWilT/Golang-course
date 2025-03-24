package storage

import (
	"database/sql"
	"github.com/AndresFWilT/postgres-db/package/invoiceitem"
)

const (
	// ...
	psqlCreateInvoiceItem = `INSERT INTO invoice_items(invoice_header_id, product_id) VALUES($1, $2) RETURNING id, created_at`
)

// PsqlInvoiceItem used for work with postgres - invoiceItem
type PsqlInvoiceItem struct {
	db *sql.DB
}

// NewPsqlInvoiceItem return a new pointer of PsqlInvoiceItem
func NewPsqlInvoiceItem(db *sql.DB) *PsqlInvoiceItem {
	return &PsqlInvoiceItem{db}
}

func (p *PsqlInvoiceItem) CreateTx(tx *sql.Tx, headerID uint, ms invoiceitem.Models) error {
	stmt, err := tx.Prepare(psqlCreateInvoiceItem)
	if err != nil {
		return err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			
		}
	}(stmt)

	for _, item := range ms {
		err = stmt.QueryRow(headerID, item.ProductID).Scan(
			&item.ID,
			&item.CreatedAt,
		)
		if err != nil {
			return err
		}
	}

	return nil
}
