package controller

import (
	"koda-b6-backend1/internal/dto"
	"koda-b6-backend1/internal/model"
	"koda-b6-backend1/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateUser(ctx *gin.Context) {

	var user model.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "bad request",
			Error:   err.Error(),
		})
		return
	}

	data, err := service.CreateUser(user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "failed create user",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, dto.Response{
		Success: true,
		Message: "create user success",
		Data: []dto.UserResponse{
			{
				ID:       data.ID,
				Email:    data.Email,
				Password: data.Password,
			},
		},
	})
}

func GetUsers(ctx *gin.Context) {

	data := service.GetUsers()

	if len(data) == 0 {
		ctx.JSON(http.StatusOK, dto.Response{
			Success: true,
			Message: "no users",
			Data:    []dto.UserResponse{},
		})
		return
	}

	var result []dto.UserResponse

	for _, dt := range data {
		result = append(result, dto.UserResponse{
			ID:    dt.ID,
			Email: dt.Email,
		})
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "all users data",
		Data:    result,
	})
}

func GetUserByID(ctx *gin.Context) {

	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "bad request",
			Error:   err.Error(),
			Data:    []dto.UserResponse{},
		})
		return
	}

	data := service.GetUserByID(id)

	if data.ID == 0 {
		ctx.JSON(http.StatusNotFound, dto.Response{
			Success: false,
			Message: "All user data",
			Data:    []dto.UserResponse{},
		})
		return
	}
	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "user found",
		Data: []dto.UserResponse{
			{
				ID:    data.ID,
				Email: data.Email,
			},
		},
	})
}
