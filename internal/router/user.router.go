package router

import (
	"koda-b6-backend1/internal/controller"

	"github.com/gin-gonic/gin"
)

// Router pakai *gin.Engine → untuk mendaftarkan endpoint
func UserRoutes(r *gin.Engine) {

	users := r.Group("/users")

	users.POST("", controller.CreateUser)
	users.GET("", controller.GetUsers)
	users.GET("/:id", controller.GetUserByID)
	users.PATCH("/:id", controller.UpdateUser)
	users.DELETE("/:id", controller.DeleteUserById)
}
