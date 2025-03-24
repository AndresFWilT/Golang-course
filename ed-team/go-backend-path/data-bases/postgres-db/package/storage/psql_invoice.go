package storage

import (
	"database/sql"
	"fmt"
	"github.com/AndresFWilT/postgres-db/package/invoice"
	"github.com/AndresFWilT/postgres-db/package/invoiceheader"
	"github.com/AndresFWilT/postgres-db/package/invoiceitem"
)

type PsqlInvoice struct {
	db            *sql.DB
	storageHeader invoiceheader.Storage
	storageItems  invoiceitem.Storage
}

func NewPsqlInvoice(db *sql.DB, h invoiceheader.Storage, i invoiceitem.Storage) *PsqlInvoice {
	return &PsqlInvoice{
		db:            db,
		storageHeader: h,
		storageItems:  i,
	}
}

func (p *PsqlInvoice) Create(m *invoice.Model) error {
	tx, err := p.db.Begin()
	if err != nil {
		return err
	}

	if err := p.storageHeader.CreateTx(tx, m.Header); err != nil {
		err := tx.Rollback()
		if err != nil {
			return err
		}
		return fmt.Errorf("header: %w", err)
	}
	fmt.Printf("Factura creada con id: %d \n", m.Header.ID)

	if err := p.storageItems.CreateTx(tx, m.Header.ID, m.Items); err != nil {
		err := tx.Rollback()
		if err != nil {
			return err
		}
		return fmt.Errorf("items: %w", err)
	}
	fmt.Printf("items creados: %d \n", len(m.Items))

	return tx.Commit()
}
