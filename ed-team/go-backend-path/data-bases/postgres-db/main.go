package main

import (
	"github.com/AndresFWilT/postgres-db/package/product"
	"github.com/AndresFWilT/postgres-db/package/storage"
	"log"
)

func main() {
	storage.NewPostgresDb()
	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewProductService(storageProduct)

	/*productModel := &product.Model{
		Name:  "Lego",
		Price: 50,
	}

	if err := serviceProduct.Create(productModel); err != nil {
		log.Fatalf("Product create: %+v", err)
	}

	fmt.Printf("Product created: %+v\n", productModel)

	productsModel, err := serviceProduct.GetAll()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(productsModel)

	m, err := serviceProduct.GetByID(3)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		fmt.Println("no hay un producto con este id")
	case err != nil:
		log.Fatalf("product.GetByID: %v", err)
	default:
		fmt.Println(m)
	}

	m := &product.Model{
		ID:    2,
		Name:  "Curso testing",
		Price: 150,
	}
	err := serviceProduct.Update(m)
	if err != nil {
		log.Fatalf("product.Update: %v", err)
	}*/
	err := serviceProduct.Delete(3)
	if err != nil {
		log.Fatalf("product.Delete: %v", err)
	}
}
