package invoice

import (
	"composition/structures/pkg/customer"
	"composition/structures/pkg/invoiceItem"
)

type Invoice struct {
	country  string
	city     string
	total    float64
	customer customer.Customer // one to one composition
	items    invoiceItem.Items // one to many composition with types
}

func New(country, city string, customer customer.Customer, items invoiceItem.Items) Invoice {
	return Invoice{
		country:  country,
		city:     city,
		customer: customer,
		items:    items,
	}
}

func (i *Invoice) SetTotal() {
	i.total = i.items.Total()
}
