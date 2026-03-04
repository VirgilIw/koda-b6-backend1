package main

import (
	"fmt"
	"koda-b6-backend1/docs"
	"koda-b6-backend1/internal/router"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// @title Koda API
// @version 1.0
// @description API documentation
// @host localhost:8888
// @BasePath /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token. Example: "Bearer eyJhbGciO..."
func main() {

	r := gin.Default()

	router.Init(r)
	docs.SwaggerInfo.BasePath = "/"
	godotenv.Load()
	port := os.Getenv("PORT")
	r.Run(fmt.Sprintf("localhost:%s", port))
}
