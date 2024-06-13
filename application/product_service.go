package application

import (
	"github.com/DerivedPuma7/go-hexagonal/application/interfaces"
)

type ProductService struct {
	Persistence interfaces.ProductPersistenceInterface
}

func NewProductService(persistence interfaces.ProductPersistenceInterface) *ProductService {
	return &ProductService{Persistence: persistence}
}

var _ interfaces.ProductServiceInterface = (*ProductService)(nil)

func (s *ProductService) Get(id string) (interfaces.ProductInterface, error) {
	product, err := s.Persistence.Get(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *ProductService) Create(name string, price float64) (interfaces.ProductInterface, error) {
	product := NewProduct()
	product.Name = name
	product.Price = price

	_, err := product.IsValid()
	if err != nil {
		return &Product{}, err
	}

	result, err := s.Persistence.Save(product)
	if err != nil {
		return &Product{}, err
	}
	return result, nil
}

func (s *ProductService) Enable(product interfaces.ProductInterface) (interfaces.ProductInterface, error) {
	err := product.Enable()
	if err != nil {
		return &Product{}, err
	}
	
	result, err := s.Persistence.Save(product);
	if err != nil {
		return &Product{}, err
	}
	return result, nil
}

func (s *ProductService) Disable(product interfaces.ProductInterface) (interfaces.ProductInterface, error) {
	err := product.Disable()
	if err != nil {
		return &Product{}, err
	}
	
	result, err := s.Persistence.Save(product);
	if err != nil {
		return &Product{}, err
	}
	return result, nil
}