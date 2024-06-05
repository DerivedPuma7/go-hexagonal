package application

import (
	"github.com/DerivedPuma7/go-hexagonal/application/interfaces"
	
)

type ProductService struct {
	Persistence interfaces.ProductPersistenceInterface
}