package main

import (
	"fmt"
	"koda-b6-backend1/docs"
	"koda-b6-backend1/internal/router"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// @title Koda-B6-Backend1
// @version         1.0
// @description     Minitask koda-b6
// @host      localhost:8888
// @BasePath  /
func main() {

	r := gin.Default()

	router.Init(r)
	docs.SwaggerInfo.BasePath = "/"
	godotenv.Load()
	port := os.Getenv("PORT")
	r.Run(fmt.Sprintf("localhost:%s", port))
}
