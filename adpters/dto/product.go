package dto

import (
	"errors"

	"github.com/matheusr42/go-hexagonal/application"
)

type Product struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	ProductStatus
}

type ProductStatus struct {
	Status string `json:"status"`
}

func NewProduct() *Product {
	return &Product{}
}

func (p *ProductStatus) Bind(status *application.Status) (*application.Status, error) {
	sts, err := MapStatus(p.Status)
	if err != nil {
		return nil, err
	}

	status = sts

	return status, nil
}

func (p *Product) Bind(product *application.Product) (*application.Product, error) {
	if p.ID != "" {
		product.ID = p.ID
	}
	product.Name = p.Name
	product.Price = p.Price
	status, err := MapStatus(p.Status)
	if err != nil {
		return nil, err
	}

	product.Status = status

	if ok, err := product.IsValid(); err != nil || !ok {
		return nil, err
	}

	return product, nil
}

func MapStatus(status string) (*application.Status, error) {
	var sts application.Status
	switch status {
	case "active":
		sts = application.ENABLED
		break
	case "inactive":
		sts = application.ENABLED
		break
	default:
		return nil, errors.New("status is required")
	}

	return &sts, nil
}
