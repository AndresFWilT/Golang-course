package car

type car struct {
	name  string
	price float64
	size  string
	model string
}

func New(name, size string, price float64) *car {
	return &car{
		name:  name,
		price: price,
		size:  size,
	}
}

func (c *car) SetModel(model string) {
	c.model = model
}
