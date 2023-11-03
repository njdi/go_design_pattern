package builder

type Product struct {
	name     string
	category string
	price    float64
}

type ProductBuilder struct {
	product *Product
}

func (pb *ProductBuilder) WithName(name string) *ProductBuilder {
	pb.product.name = name
	return pb
}

func (pb *ProductBuilder) WithCategory(category string) *ProductBuilder {
	pb.product.category = category
	return pb
}

func (pb *ProductBuilder) WithPrice(price float64) *ProductBuilder {
	pb.product.price = price
	return pb
}

func (pb *ProductBuilder) Build() *Product {
	return pb.product
}

func NewProductBuilder() *ProductBuilder {
	return &ProductBuilder{&Product{}}
}
