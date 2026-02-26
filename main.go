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

	// get all user
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
	// detail user
	r.GET("/users/:id", func(ctx *gin.Context) {
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

		for _, u := range users {
			if u.ID == id {
				ctx.JSON(http.StatusOK, Response{
					Success: true,
					Message: "user found",
					Data:    []DataResponse{u},
				})
				return
			}
		}

		ctx.JSON(http.StatusNotFound, Response{
			Success: false,
			Message: "user not found",
			Error:   "id does not exist",
		})
	})

	// add user
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

	// edit user
	r.PATCH("/users/:id", func(ctx *gin.Context) {
		id, _ := strconv.Atoi(ctx.Param("id"))

		var data DataResponse
		if err := ctx.ShouldBindJSON(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, Response{
				Success: false,
				Message: "bad request",
				Error:   err.Error(),
			})
			return
		}

		for i := range users {
			if users[i].ID == id {

				// cek email duplicate
				if data.Email != "" {
					for _, u := range users {
						if u.Email == data.Email && u.ID != id {
							ctx.JSON(http.StatusBadRequest, Response{
								Success: false,
								Message: "email already used",
							})
							return
						}
					}
					users[i].Email = data.Email
				}

				if data.Password != "" {
					users[i].Password = data.Password
				}

				ctx.JSON(http.StatusOK, Response{
					Success: true,
					Message: "user updated",
					Data:    users,
				})
				return
			}
		}

		ctx.JSON(http.StatusNotFound, Response{
			Success: false,
			Message: "user not found",
		})
	})

	// delete user
	r.DELETE("/users/:id", func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, _ := strconv.Atoi(idParam)
		var result []DataResponse
		found := false

		for _, v := range users {
			if v.ID == id {
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

	r.POST("/register", func(ctx *gin.Context) {
		var data DataResponse
		if err := ctx.ShouldBindJSON(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, Response{
				Success: false,
				Message: "bad request",
				Error:   err.Error(),
				Data:    []DataResponse{},
			})
			return
		}

		if len(data.Password) < 8 {
			ctx.JSON(http.StatusBadRequest, Response{
				Success: false,
				Message: "bad request",
				Error:   "password must be at least 8 characters",
			})
			return
		}

		for _, v := range users {
			if v.Email == data.Email {
				ctx.JSON(http.StatusBadRequest, Response{
					Success: false,
					Message: "bad request",
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
			Message: "success register account",
			Data:    users,
		})
	})

	r.POST("/login", func(ctx *gin.Context) {
		var data DataResponse

		if err := ctx.ShouldBindJSON(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, Response{
				Success: false,
				Message: "bad request",
				Error:   err.Error(),
			})
			return
		}

		for _, u := range users {
			if u.Email == data.Email && u.Password == data.Password {
				ctx.JSON(http.StatusOK, Response{
					Success: true,
					Message: "login success",
				})
				return
			}
		}

		ctx.JSON(http.StatusUnauthorized, Response{
			Success: false,
			Message: "login failed",
			Error:   "email or password incorrect",
		})
	})

	r.Run("localhost:8888")
}
