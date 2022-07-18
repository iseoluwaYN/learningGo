package cart

type Category string

type Product struct {
	Id          int
	Name        string
	Price       float64
	Description string
	Category    Category
}

type Item struct {
	Quantity int
	Product  Product
}
