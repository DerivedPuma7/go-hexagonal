package application

import (
	"github.com/DerivedPuma7/go-hexagonal/application/interfaces"
)

type ProductService struct {
	Persistence interfaces.ProductPersistenceInterface
}

var _ interfaces.ProductServiceInterface = (*ProductService)(nil)

func (s *ProductService) Get(id string) (interfaces.ProductInterface, error) {

}

// func (s *ProductService) Create(name string, price float64) (interfaces.ProductInterface, error) {

// }

// func (s *ProductService) Enable(product interfaces.ProductInterface) (interfaces.ProductInterface, error) {

// }

// func (s *ProductService) Disable(product interfaces.ProductInterface) (interfaces.ProductInterface, error) {

// }