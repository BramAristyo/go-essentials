package testing

type Product struct {
	ID    int
	Name  string
	Price int
}

func CalculateTotal(products []Product) int {
	total := 0

	for _, p := range products {
		total += p.Price
	}

	return total
}
