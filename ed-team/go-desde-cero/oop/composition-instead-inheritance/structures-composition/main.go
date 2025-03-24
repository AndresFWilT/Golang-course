package main

import (
	"composition/structures/pkg/customer"
	"composition/structures/pkg/invoice"
	"composition/structures/pkg/invoiceItem"
	"fmt"
)

func main() {
	customerInstance := customer.New("Andres", "Calle 123", "3001234")
	invoiceItemInstance := invoiceItem.NewItems(invoiceItem.New(1, "Bed", 1234), invoiceItem.New(2, "Library", 12342), invoiceItem.New(3, "Desktop", 12349))
	invoiceInstance := invoice.New("Colombia", "Bogota", customerInstance, invoiceItemInstance)
	invoiceInstance.SetTotal()
	fmt.Printf("%+v", invoiceInstance)
}
