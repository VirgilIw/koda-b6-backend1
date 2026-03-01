package router

import (
	"koda-b6-backend1/internal/controller"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.Engine) {

	products := r.Group("/products")

	products.GET("", controller.GetAllProduct)
	products.GET("/:id", controller.GetProductById)
	products.POST("", controller.CreateProduct)
	products.PATCH("/:id", controller.EditProductById)
	products.DELETE("/:id", controller.DeleteProductById)
}
