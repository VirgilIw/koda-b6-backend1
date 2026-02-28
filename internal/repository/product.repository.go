package repository

import "koda-b6-backend1/internal/model"

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
