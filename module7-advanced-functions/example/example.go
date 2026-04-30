package example

type Product struct {
	Name    string
	Price   float64
	InStock bool
}

func Filter(products []Product, filterFunc func(p Product) bool) []Product {
	var result []Product
	for _, product := range products {
		if filterFunc(product) {
			result = append(result, product)
		}
	}
	return result
}
