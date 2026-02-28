package controller

import (
	"koda-b6-backend1/internal/dto"
	"koda-b6-backend1/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateProduct(ctx *gin.Context) {
	var req dto.ProductRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ResponseProduct{
			Success: false,
			Message: "bad request",
			Error:   err.Error(),
			Data:    []dto.ProductResponse{},
		})
		return
	}

	data, err := service.CreateProduct(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ResponseProduct{
			Success: false,
			Message: "failed create product",
			Error:   err.Error(),
			Data:    []dto.ProductResponse{},
		})
		return
	}

	ctx.JSON(http.StatusCreated, dto.ResponseProduct{
		Success: true,
		Message: "Success create product",
		Data:    []dto.ProductResponse{data},
	})
}

func GetAllProduct(ctx *gin.Context) {
	data := service.GetAllProduct()
	if len(data) == 0 {
		ctx.JSON(http.StatusBadRequest, dto.ResponseProduct{
			Success: false,
			Message: "bad request",
			Error:   "products data empty",
			Data:    []dto.ProductResponse{},
		})
		return
	}
	ctx.JSON(http.StatusOK, dto.ResponseProduct{
		Success: true,
		Message: "success get all data",
		Data:    data,
	})
}
