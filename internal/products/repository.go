package products

import "errors"

type Repository interface {
	GetAllBySeller(sellerID string) ([]Product, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAllBySeller(sellerID string) ([]Product, error) {
	var prodList []Product
	prodList = append(prodList, Product{
		ID:          "mock",
		SellerID:    "FEX112AC",
		Description: "generic product",
		Price:       123.55,
	})
	var response []Product
	for i, product := range prodList {
		if product.SellerID == sellerID {
			response = append(response, prodList[i])
		}
	}
	if len(response) == 0 {
		return nil, errors.New("Error en el repositorio")
	}
	return response, nil
}
