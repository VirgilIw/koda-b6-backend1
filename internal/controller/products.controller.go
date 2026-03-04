package controller

import (
	"koda-b6-backend1/internal/dto"
	"koda-b6-backend1/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Accept dipakai kalau request body ada.

// CreateProduct godoc
// @Summary      Create a product
// @Description  Create new product
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        product body dto.ProductRequest true "Create product"
// @Success      201  {object}  dto.ResponseProduct
// @Failure      400  {object}  dto.ResponseProduct
// @Failure      500  {object}  dto.ResponseProduct
// @Router       /products [post]
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

// GetAllProduct godoc
// @Summary      Get all products
// @Description  Get all products
// @Tags         products
// @Produce      json
// @Success      200  {object}  dto.ResponseProduct
// @Failure      500  {object}  dto.ResponseProduct
// @Router       /products [get]
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

// GetProductById godoc
// @Summary      Get Product By Id
// @Description  Get Product By Id
// @Tags         products
// @Produce      json
// @Param        id   path   int   true   "Product ID"
// @Success      200  {object}  dto.ResponseProduct
// @Failure      400  {object}  dto.ResponseProduct
// @Failure      404  {object}  dto.ResponseProduct
// @Failure      500  {object}  dto.ResponseProduct
// @Router       /products/{id} [get]
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

// EditProductById godoc
// @Summary      Edit Product By Id
// @Description  Update product by ID
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id path int true "Product ID"
// @Param        product body dto.ProductRequest true "Product data"
// @Success      200  {object}  dto.ResponseProduct
// @Failure      400  {object}  dto.ResponseProduct
// @Failure      404  {object}  dto.ResponseProduct
// @Failure      500  {object}  dto.ResponseProduct
// @Router       /products/{id} [patch]
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

// DeleteProduct godoc
// @Summary      Delete a product
// @Description  Delete a product
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id path int true "Delete product"
// @Success      200  {object}  dto.ResponseProduct
// @Failure      400  {object}  dto.ResponseProduct
// @Failure      500  {object}  dto.ResponseProduct
// @Router       /products/{id} [delete]
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
