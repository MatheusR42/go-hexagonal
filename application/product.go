package application

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

type ProductInterface interface {
	IsValid() (bool, error)
	Enable() error
	Disable() error
	GetId() string
	GetName() string
	GetStatus() string
	GetPrice() float64
}

type Status int

const (
	DISABLED Status = iota
	ENABLED
)

func Init() {
	validate = validator.New()

	// validateStruct()
	// validateVariable()
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

func (p *Product) GetId() string {
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
