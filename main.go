package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type DataResponse struct {
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required"`
}

type Response struct {
	Success bool           `json:"success"`
	Message string         `json:"message"`
	Error   string         `json:"error,omitempty"`
	Data    []DataResponse `json:"data,omitempty"`
}

// omitempty = field kosong, jangan tampilkan di JSON response.

var users []DataResponse

func main() {
	r := gin.Default()

	r.GET("/users", func(ctx *gin.Context) {
		if len(users) == 0 {
			ctx.JSON(http.StatusNotFound, Response{
				Success: false,
				Message: "no users found",
				Error:   "data is empty",
			})
			return
		}

		ctx.JSON(http.StatusOK, Response{
			Success: true,
			Message: "get user data success",
			Data:    users,
		})
	})

	r.POST("/users", func(ctx *gin.Context) {
		var data DataResponse

		if err := ctx.ShouldBindJSON(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, Response{
				Success: false,
				Message: "Bad Request",
				Error:   err.Error(),
			})
			return
		}

		for _, u := range users {
			if u.Email == data.Email {
				ctx.JSON(http.StatusBadRequest, Response{
					Success: false,
					Message: "Bad Request",
					Error:   "email already registered",
				})
				return
			}
		}

		users = append(users, data)

		ctx.JSON(http.StatusCreated, Response{
			Success: true,
			Message: "success create user",
			Data:    users,
		})
	})

	r.Run("localhost:8888")
}
