package repository

import (
	"errors"
	"koda-b6-backend1/internal/dto"
	"koda-b6-backend1/internal/model"
)

var products []model.Product
var IdProduct int

func CreateProduct(product model.Product) model.Product {
	IdProduct++
	product.ID = IdProduct
	products = append(products, product)

	return product
}

func FindAllProduct() []model.Product {
	return products
}

func GetProductById(id int) (model.Product, error) {
	for _, p := range products {
		if p.ID == id {
			return p, nil
		}
	}

	return model.Product{}, errors.New("product not found")
}

func EditProductById(id int, req dto.ProductRequest) error {

	for i, p := range products {

		if p.ID == id {

			products[i].Name = req.Name
			products[i].Description = req.Description
			products[i].Rating = req.Rating
			products[i].Stock = req.Stock
			products[i].Images = req.Images

			return nil
		}
	}

	return errors.New("product not found")
}

func DeleteProductById(id int) error {

	var newProducts []model.Product
	found := false

	for _, p := range products {
		if p.ID == id {
			found = true
			continue
		}

		newProducts = append(newProducts, p)
	}

	products = newProducts

	if !found {
		return errors.New("product not found")
	}

	return nil
}
