package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Init(app *gin.Engine) {
	UserRoutes(app)
	ProductRoutes(app)

	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
