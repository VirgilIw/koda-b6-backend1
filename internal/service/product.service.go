package service

import (
	"errors"
	"koda-b6-backend1/internal/dto"
	"koda-b6-backend1/internal/model"
	"koda-b6-backend1/internal/repository"
)

func CreateProduct(req *dto.ProductRequest) (dto.ProductResponse, error) {

	products := repository.FindAllProduct()

	for _, p := range products {
		if p.Name == req.Name {
			return dto.ProductResponse{}, errors.New("product already created")
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
