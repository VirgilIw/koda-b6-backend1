package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/matthewhartstonge/argon2"
)

type DataResponse struct {
	ID       int    `json:"id"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password,omitempty" form:"password" binding:"required"`
}

type DataResponse struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Response struct {
	Success bool           `json:"success"`
	Message string         `json:"message"`
	Error   string         `json:"error,omitempty"`
	Data    []DataResponse `json:"results,omitempty"`
}

var users []DataResponse
var id int

func main() {
	r := gin.Default()

	r.POST("/users", func(ctx *gin.Context) {

		var data DataResponse

		if err := ctx.ShouldBindJSON(&data); err != nil {
			ctx.JSON(http.StatusInternalServerError, Response{
				Success: false,
				Message: "internal server error",
				Error:   "server error",
				Data:    []DataResponse{},
			})
		} else {
			for _, i := range users {
				if i.Email == data.Email {
					ctx.JSON(http.StatusBadRequest, Response{
						Success: false,
						Message: "bad request",
						Error:   "email already registered",
					})
					return
				}
			}
			id++
			data.Id = id
			users = append(users, data)
			ctx.JSON(http.StatusCreated, Response{
				Success: true,
				Message: "success create user",
				Data:    users,
			})
		}
	})

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
		ctx.JSON(http.StatusOK, Response{
			Success: true,
			Message: "success get user data",
			Data:    users,
		})
	})

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
		for _, u := range users {
			if u.Id == id {
				ctx.JSON(http.StatusOK, Response{
					Success: true,
					Message: "user found, id:" + " " + strconv.Itoa(u.Id),
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
		// hashed password
		encoded, err := argon.HashEncoded([]byte(data.Password))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, Response{
				Success: false,
				Message: "failed to hash password",
				Error:   err.Error(),
			})
			return
		}

		// masukin hashpassword ke variable
		data.Password = string(encoded)
		lastID++
		data.ID = lastID
		users = append(users, data)
		fmt.Println(data.Password)
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
			// cek email sama gak antara di tempat simpan dengan yang di input user
			if u.Email == data.Email {
				// bandingkan hash dari user dengan hash dari tempat penyimpanan :)
				ok, err := argon2.VerifyEncoded(
					[]byte(data.Password),
					[]byte(u.Password),
				)
				if err != nil || !ok {
					ctx.JSON(http.StatusUnauthorized, Response{
						Success: false,
						Message: "login failed",
						Error:   "email or password incorrect",
					})
					return
				}

				ctx.JSON(http.StatusOK, Response{
					Success: true,
					Message: "login success",
					Data: []DataResponse{
						{
							ID:    u.ID,
							Email: u.Email,
						},
					},
				})
				return
			}
		}
	})

	r.Run("localhost:8888")
	})
}