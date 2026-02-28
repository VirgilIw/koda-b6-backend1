package router

import "github.com/gin-gonic/gin"

func Init(app *gin.Engine) {
	UserRoutes(app)
	ProductRoutes(app)
}
