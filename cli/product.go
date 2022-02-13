package cli

import (
	"fmt"

	"github.com/matheusr42/go-hexagonal/application"
)

type Action string

const (
	CREATE  Action = "CREATE"
	ENABLE         = "ENABLE"
	DISABLE        = "DISABLE"
	GET            = "GET"
)

func Run(service application.ProductServiceInterface, action Action, productID string, productName string, price float64) (string, error) {
	result := ""

	switch action {
	case CREATE:
		_, err := service.Create(productName, price)
		if err != nil {
			return result, err
		}
		result = "product created"
	case ENABLE:
		product, err := service.Get(productID)
		if err != nil {
			return result, err
		}

		product, err = service.Enable(product)
		if err != nil {
			return result, err
		}

		result = "product enabled"
	case DISABLE:
		product, err := service.Get(productID)
		if err != nil {
			return result, err
		}

		product, err = service.Disable(product)
		if err != nil {
			return result, err
		}

		result = "product disabled"
	case GET:
		product, err := service.Get(productID)
		if err != nil {
			return result, err
		}

		result = fmt.Sprintf("product: %s", product.GetName())
	default:
		return "", fmt.Errorf("invalid action")
	}

	return result, nil
}
