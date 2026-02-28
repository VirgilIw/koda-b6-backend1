package main

import (
	"koda-b6-backend1/internal/router"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	router.Init(r)

	r.Run(":8080")
}
