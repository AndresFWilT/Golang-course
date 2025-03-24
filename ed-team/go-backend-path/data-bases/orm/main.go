package main

import (
	"github.com/AndresFWilT/postgres-orm/model"
	"github.com/AndresFWilT/postgres-orm/storage"
)

func main() {
	driver := storage.Postgres
	storage.New(driver)

	// Create product
	/*product1 := model.Product{
		Name:  "Curso de Go",
		Price: 120,
	}

	obs := "testing with Go"
	product2 := model.Product{
		Name:         "Curso de Testing",
		Price:        150,
		Observations: &obs,
	}

	product3 := model.Product{
		Name:  "Curso de Python",
		Price: 200,
	}

	storage.DB().Create(&product1)
	storage.DB().Create(&product2)
	storage.DB().Create(&product3)*/

	// Get all products

	/*products := make([]model.Product, 0)
	storage.DB().Find(&products)

	for _, product := range products {
		fmt.Printf("%d - %s\n", product.ID, product.Name)
	}*/

	// Get a Product

	/*myProduct := model.Product{}

	storage.DB().First(&myProduct, 2)
	fmt.Println(myProduct)*/

	// Update a Product

	/*myProduct := model.Product{}
	myProduct.ID = 2

	storage.DB().Model(&myProduct).Updates(
		model.Product{Name: "Curso de Java", Price: 120},
	)*/
	/*
		// soft Delete
		myProduct := model.Product{}
		myProduct.ID = 4

		storage.DB().Delete(&myProduct)

		// Delete Permanent
		myProduct1 := model.Product{}
		myProduct1.ID = 6

		storage.DB().Unscoped().Delete(&myProduct1)*/

	// Transactions } Associations
	// Create foreign keys only in GORM not in the database !IMPORTANT
	storage.DB().Model(&model.InvoiceItem{}).AddForeignKey("product_id", "products(id)", "RESTRICT", "RESTRICT")

	storage.DB().Model(&model.InvoiceItem{}).AddForeignKey("invoice_header_id", "invoice_headers(id)", "RESTRICT", "RESTRICT")

	invoice := model.InvoiceHeader{
		Client: "Andres Felipe",
		InvoiceItems: []model.InvoiceItem{
			{ProductID: 1},
		},
	}

	storage.DB().Create(&invoice)
}
