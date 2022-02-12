package application

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

type ProductInterface interface {
	IsValid() (bool, error)
	Enable() error
	Disable() error
	GetID() string
	GetName() string
	GetStatus() Status
	GetPrice() float64
}

type ProductServiceInterface interface {
	Get(ID string) (ProductInterface, error)
	Create(name string, price float64) (ProductInterface, error)
	Enable(Product ProductInterface) (ProductInterface, error)
	Disable(Product ProductInterface) (ProductInterface, error)
}

type ProductReaderInterface interface {
	Get(ID string) (ProductInterface, error)
}

type ProductWriterInterface interface {
	Save(product ProductInterface) (ProductInterface, error)
}

type ProductPersistenceInterface interface {
	ProductReaderInterface
	ProductWriterInterface
}

type Status int

const (
	DISABLED Status = iota
	ENABLED
)

func Init() {
	validate = validator.New()
}

func (s Status) String() string {
	return [...]string{"DISABLED", "ENABLED"}[s]
}

type Product struct {
	ID     string  `validate:"uuid4"`
	Name   string  `validate:"required"`
	Price  float64 `validate:"gt=-1"`
	Status *Status `validate:"required"`
}

func NewProduct() *Product {
	status := DISABLED
	product := Product{
		ID:     uuid.New().String(),
		Status: &status,
	}

	return &product
}

func (p *Product) IsValid() (bool, error) {
	err := validate.Struct(p)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (p *Product) Enable() error {
	if p.Price > 0 {
		status := ENABLED
		p.Status = &status
		return nil
	}

	return errors.New("the price must be greater than zero")
}

func (p *Product) Disable() error {
	if p.Price == 0 {
		status := DISABLED
		p.Status = &status
		return nil
	}

	return errors.New("the price must be zero to disable product")
}

func (p *Product) GetID() string {
	return p.ID
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) GetStatus() Status {
	return *p.Status
}

func (p *Product) GetPrice() float64 {
	return p.Price
}
