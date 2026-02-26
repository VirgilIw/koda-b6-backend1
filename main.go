package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DataResponse struct {
	ID       int    `json:"id"`
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
var lastID int

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
		lastID++
		data.ID = lastID
		users = append(users, data)

		ctx.JSON(http.StatusCreated, Response{
			Success: true,
			Message: "success create user",
			Data:    users,
		})
	})

	r.PATCH("/users/:id", func(ctx *gin.Context) {
		idParam := ctx.Param("id")

		id, err := strconv.Atoi(idParam)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, Response{
				Success: false,
				Message: "invalid id",
				Error:   err.Error(),
			})
			return
		}

		var data DataResponse

		if err := ctx.ShouldBindJSON(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, Response{
				Success: false,
				Message: "bad request",
				Error:   err.Error(),
			})
			return
		}

		found := false

		for i, u := range users {
			if u.ID == id {
				found = true

				if data.Email != "" {
					users[i].Email = data.Email
				}

				if data.Password != "" {
					users[i].Password = data.Password
				}

				break
			}
		}

		if !found {
			ctx.JSON(http.StatusNotFound, Response{
				Success: false,
				Message: "user not found",
				Error:   "id does not exist",
			})
			return
		}

		ctx.JSON(http.StatusOK, Response{
			Success: true,
			Message: "user updated",
			Data:    users,
		})
	})

	r.DELETE("/users/:id", func(ctx *gin.Context) {
		idParam := ctx.Param("id")

		var result []DataResponse
		found := false

		for _, v := range users {
			if strconv.Itoa(v.ID) == idParam {
				found = true
				continue
			}
			result = append(result, v)
		}

		if !found {
			ctx.JSON(http.StatusNotFound, Response{
				Success: false,
				Message: "user not found",
				Error:   "id does not exist",
			})
			return
		}

		users = result

		ctx.JSON(http.StatusOK, Response{
			Success: true,
			Message: "success delete user",
			Data:    result,
		})
	})
	r.Run("localhost:8888")
}
