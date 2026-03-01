package controller

import (
	"koda-b6-backend1/internal/dto"
	"koda-b6-backend1/internal/service"
	"net/http"
	"strconv"

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

func GetProductById(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)

	if err != nil || id <= 0 {
		ctx.JSON(http.StatusBadRequest, dto.ResponseProduct{
			Success: false,
			Message: "bad request",
			Error:   "invalid product id",
			Data:    []dto.ProductResponse{},
		})
		return
	}

	data, err := service.GetProductById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, dto.ResponseProduct{
			Success: false,
			Message: "bad request",
			Error:   err.Error(),
			Data:    []dto.ProductResponse{},
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.ResponseProduct{
		Success: true,
		Message: "success get product",
		Data:    []dto.ProductResponse{data},
	})
}

func EditProductById(ctx *gin.Context) {

	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil || id <= 0 {
		ctx.JSON(http.StatusBadRequest, dto.ResponseProduct{
			Success: false,
			Message: "bad request",
			Error:   "invalid product id",
			Data:    []dto.ProductResponse{},
		})
		return
	}

	var req dto.ProductRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ResponseProduct{
			Success: false,
			Message: "invalid request body",
			Error:   err.Error(),
			Data:    []dto.ProductResponse{},
		})
		return
	}

	err = service.EditProductById(id, req)
	if err != nil {
		ctx.JSON(http.StatusNotFound, dto.ResponseProduct{
			Success: false,
			Message: "product not found",
			Error:   err.Error(),
			Data:    []dto.ProductResponse{},
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.ResponseProduct{
		Success: true,
		Message: "success edit product",
		Data:    []dto.ProductResponse{},
	})
}

func DeleteProductById(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)

	if err != nil || id <= 0 {
		ctx.JSON(http.StatusBadRequest, dto.ResponseProduct{
			Success: false,
			Message: "bad request",
			Error:   "invalid product id",
			Data:    []dto.ProductResponse{},
		})
		return
	}

	err = service.DeleteProductById(id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, dto.ResponseProduct{
			Success: false,
			Message: "product not found",
			Error:   err.Error(),
			Data:    []dto.ProductResponse{},
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.ResponseProduct{
		Success: true,
		Message: "product deleted",
		Data:    []dto.ProductResponse{},
	})
}
