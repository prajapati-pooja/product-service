package internal

type Service interface {
	GetAll(productType string) []Product
}

type productService struct {
	repository Repository
}

func NewProductService(productRepository Repository) Service {
	return productService{repository: productRepository}
}

func (ps productService) GetAll(productType string) []Product {
	products := ps.repository.Get()
	var specificProducts []Product
	for _, product := range products {
		if product.ProductType == productType {
			specificProducts = append(specificProducts, product)
		}
	}
	return specificProducts
}
