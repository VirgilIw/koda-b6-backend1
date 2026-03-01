package service

import (
	"errors"
	"koda-b6-backend1/internal/dto"
	"koda-b6-backend1/internal/model"
	"koda-b6-backend1/internal/repository"
)

var ErrProductAlreadyExists = errors.New("product already exists")

func CreateProduct(req *dto.ProductRequest) (dto.ProductResponse, error) {

	products := repository.FindAllProduct()

	for _, p := range products {
		if p.Name == req.Name {
			return dto.ProductResponse{}, ErrProductAlreadyExists
		}
	}

	product := model.Product{
		Name:        req.Name,
		Description: req.Description,
		Rating:      req.Rating,
		Stock:       req.Stock,
		Images:      req.Images,
	}

	data := repository.CreateProduct(product)

	return dto.ProductResponse{
		Id:          data.ID,
		Name:        data.Name,
		Description: data.Description,
		Rating:      data.Rating,
		Stock:       data.Stock,
		Images:      data.Images,
	}, nil
}

func GetAllProduct() []dto.ProductResponse {
	data := repository.FindAllProduct()

	var result []dto.ProductResponse

	for _, p := range data {
		result = append(result, dto.ProductResponse{
			Id:          p.ID,
			Name:        p.Name,
			Description: p.Description,
			Rating:      p.Rating,
			Stock:       p.Stock,
			Images:      p.Images,
		})
	}
	return result
}

func GetProductById(id int) (dto.ProductResponse, error) {
	data, err := repository.GetProductById(id)

	if err != nil || data.ID == 0 {
		return dto.ProductResponse{}, err
	}

	return dto.ProductResponse{
		Id:          data.ID,
		Name:        data.Name,
		Description: data.Description,
		Rating:      data.Rating,
		Stock:       data.Stock,
		Images:      data.Images,
	}, nil
}

func EditProductById(id int, req dto.ProductRequest) error {
	return repository.EditProductById(id, req)
}

func DeleteProductById(id int) error {
	return repository.DeleteProductById(id)
}
