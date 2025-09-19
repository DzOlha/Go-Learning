package product

type Product struct {
	Title string
	ID    int64
	Price float64
}

func New(title string, id int64, price float64) *Product {
	return &Product{
		Title: title,
		ID:    id,
		Price: price,
	}
}
