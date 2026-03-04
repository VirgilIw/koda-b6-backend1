package controller

import (
	"koda-b6-backend1/internal/dto"
	"koda-b6-backend1/internal/model"
	"koda-b6-backend1/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateUser godoc
// @Summary      Create a user
// @Description  Create new user
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        user body model.User true "Create user"
// @Success      201  {object}  dto.UserResponse
// @Failure      400  {object}  dto.UserResponse
// @Failure      500  {object}  dto.UserResponse
// @Router       /users [post]
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
				Email:    data.Email,
				Password: data.Password,
			},
		},
	})
}

// GetUsers godoc
// @Summary      Get users
// @Description  Get users
// @Tags         Users
// @Produce      json
// @Success      200  {object}  dto.UserResponse
// @Failure      400  {object}  dto.UserResponse
// @Failure      500  {object}  dto.UserResponse
// @Router       /users [get]
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
			Email: dt.Email,
		})
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "all users data",
		Data:    result,
	})
}

// GetUsersById godoc
// @Summary      Get users
// @Description  Get users
// @Tags         Users
// @Produce      json
// @Param        id   path   int   true   "User ID"
// @Success      200  {object}  dto.UserResponse
// @Failure      400  {object}  dto.UserResponse
// @Failure      500  {object}  dto.UserResponse
// @Router       /users/{id} [get]
func GetUserByID(ctx *gin.Context) {

	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "bad request",
			Error:   err.Error(),
		})
		return
	}

	data := service.GetUserByID(id)

	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "user found",
		Data: []dto.UserResponse{
			{
				Email: data.Email,
			},
		},
	})
}

// UpdateUser godoc
// @Summary      Update user By Id
// @Description  Update user by ID
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        id path int true "User ID"
// @Param        user body dto.UpdateUserRequest true "User data"
// @Success      200  {object}  dto.UserResponse
// @Failure      400  {object}  dto.UserResponse
// @Failure      404  {object}  dto.UserResponse
// @Failure      500  {object}  dto.UserResponse
// @Router       /users/{id} [patch]
func UpdateUser(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "bad request",
			Error:   err.Error(),
		})
		return
	}
	var req dto.UpdateUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "bad request",
			Error:   err.Error(),
			Data:    []dto.UserResponse{},
		})
	}

	data, err := service.UpdateUser(id, req)

	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "update data success",
		Data: []dto.UserResponse{
			{
				Email: data.Email,
			},
		},
	})

}

// DeleteUser godoc
// @Summary      Delete user By Id
// @Description  Delete user by ID
// @Tags         Users
// @Produce      json
// @Param        id path int true "Delete User ID"
// @Success      200  {object}  dto.UserResponse
// @Failure      400  {object}  dto.UserResponse
// @Failure      404  {object}  dto.UserResponse
// @Failure      500  {object}  dto.UserResponse
// @Router       /users/{id} [delete]
func DeleteUserById(ctx *gin.Context) {

	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "bad request",
			Error:   "invalid user id",
			Data:    []dto.UserResponse{},
		})
		return
	}

	err = service.DeleteUserById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, dto.Response{
			Success: false,
			Message: "user not found",
			Error:   err.Error(),
			Data:    []dto.UserResponse{},
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "success delete user",
		Data:    []dto.UserResponse{},
	})
}
