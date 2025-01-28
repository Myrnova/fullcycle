package cli

import (
	"fmt"
	"myrnova/hexagonal/application"
)

func Run(service application.ProductServiceInterface, action string, productId string, productName string, price float64) (string, error) {
	var result string

	switch action {
	case "create":
		product, err := service.Create(productName, price)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product ID #%s with name %s has been created with price %.2f", product.GetId(), productName, price)
	case "enable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		_, err = service.Enable(product)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product #%s has been enabled", product.GetName())
	case "disable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		_, err = service.Disable(product)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product #%s has been disabled", product.GetName())

	default:
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product ID %s\nName %s\nPrice %.2f\nStatus %s",
			product.GetId(), product.GetName(), product.GetPrice(), product.GetStatus())
	}
	return result, nil
}
