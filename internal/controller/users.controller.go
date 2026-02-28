package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matthewhartstonge/argon2"
)

var argon = argon2.DefaultConfig()

type User struct {
	ID       int
	Email    string
	Password string
}

type CreateUserResponse struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}

type Response struct {
	Success bool                 `json:"success"`
	Message string               `json:"message"`
	Error   string               `json:"error,omitempty"`
	Data    []CreateUserResponse `json:"data,omitempty"`
}

var users []User
var lastID int

// CREATE USER
func CreateUser(ctx *gin.Context) {

	var user User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: "bad request",
			Error:   err.Error(),
		})
		return
	}

	if len(user.Password) < 8 {
		ctx.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: "bad request",
			Error:   "password must be at least 8 characters",
		})
		return
	}

	for _, u := range users {
		if u.Email == user.Email {
			ctx.JSON(http.StatusBadRequest, Response{
				Success: false,
				Message: "email already registered",
			})
			return
		}
	}

	encoded, err := argon.HashEncoded([]byte(user.Password))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Response{
			Success: false,
			Message: "failed hash password",
			Error:   err.Error(),
		})
		return
	}

	lastID++
	user.ID = lastID
	user.Password = string(encoded)

	users = append(users, user)

	ctx.JSON(http.StatusCreated, Response{
		Success: true,
		Message: "create user success",
		Data: []CreateUserResponse{
			{
				ID:       user.ID,
				Email:    user.Email,
				Password: user.Password,
			},
		},
	})
}

// GET ALL USERS
func GetUsers(ctx *gin.Context) {
	data := users
	if len(data) == 0 {
		ctx.JSON(http.StatusOK, Response{
			Success: true,
			Message: "no users",
			Data:    []CreateUserResponse{},
		})
		return
	}

	var user []CreateUserResponse
	for _, dt := range data {
		user = append(user, CreateUserResponse{
			ID:    dt.ID,
			Email: dt.Email,
		})
	}

	ctx.JSON(http.StatusOK, Response{
		Success: true,
		Message: "all users data",
		Data:    user,
	})
}

// GET USER BY ID
func GetUserByID(ctx *gin.Context) {

}
