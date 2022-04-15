package dto

import "github.com/matheusr42/go-hexagonal/application"

type Product struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	Status string  `json:"status"`
}

func NewProduct() *Product {
	return &Product{}
}

func (p *Product) Bind(product *application.Product) (*application.Product, error) {
	if p.ID != "" {
		product.ID = p.ID
	}
	product.Name = p.Name
	product.Price = p.Price
	status := MapStatus(p.Status)
	product.Status = &status

	if ok, err := product.IsValid(); err != nil || !ok {
		return nil, err
	}

	return product, nil
}

func MapStatus(status string) application.Status {
	switch status {
	case "active":
		return application.ENABLED
	case "inactive":
		return application.DISABLED
	default:
		return application.DISABLED
	}
}
