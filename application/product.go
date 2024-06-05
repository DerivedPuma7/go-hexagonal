package application

import (
	"errors"
	"github.com/DerivedPuma7/go-hexagonal/application/interfaces"
	
	"github.com/google/uuid"
	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

const (
	DISABLED = "disabled"
	ENABLED = "enabled"
)

var _ interfaces.ProductInterface = (*Product)(nil)

type Product struct {
	ID string `valid:"uuidv4"`
	Name string `valid:"required"`
	Status string `valid:"required"`
	Price float64 `valid:"float,optional"`
}

func NewProduct() *Product {
	product := Product {
		ID: uuid.New().String(),
		Status: DISABLED,
	}
	return &product;
}

func (p *Product) IsValid() (bool, error) {
	if p.Status == "" {
		p.Status = DISABLED
	}

	if p.Status != ENABLED && p.Status != DISABLED {
		return false, errors.New("status must be enabled or disabled")
	}

	if p.Price < 0 {
		return false, errors.New("price must be greater or equal zero")
	}

	_, err := govalidator.ValidateStruct(p)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (p *Product) Enable() error {
	if p.Price > 0 {
		p.Status = ENABLED
		return nil
	}
	return errors.New("price must be greater than zero to enable product")
}

func (p *Product) Disable() error {
	if p.Price == 0 {
		p.Status = DISABLED
		return nil
	}
	return errors.New("price must be zero in order to disable a product")
}

func (p *Product) GetID() string {
	return p.ID
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) GetStatus() string {
	return p.Status
}

func (p *Product) GetPrice() float64 {
	return p.Price
}
