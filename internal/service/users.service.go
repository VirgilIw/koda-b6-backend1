package service

import (
	"errors"
	"koda-b6-backend1/internal/dto"
	"koda-b6-backend1/internal/model"
	"koda-b6-backend1/internal/repository"

	"github.com/matthewhartstonge/argon2"
)

var argon = argon2.DefaultConfig()

func CreateUser(user model.User) (dto.UserResponse, error) {

	if len(user.Password) < 8 {
		return dto.UserResponse{}, errors.New("password must be at least 8 characters")
	}

	users := repository.FindAll()

	for _, u := range users {
		if u.Email == user.Email {
			return dto.UserResponse{}, errors.New("email already registered")
		}
	}

	hash, err := argon.HashEncoded([]byte(user.Password))
	if err != nil {
		return dto.UserResponse{}, err
	}

	user.Password = string(hash)

	data := repository.Create(user)

	return dto.UserResponse{
		ID:       data.ID,
		Email:    data.Email,
		Password: data.Password,
	}, nil
}

func GetUsers() []dto.UserResponse {

	users := repository.FindAll()

	var result []dto.UserResponse

	for _, u := range users {
		result = append(result, dto.UserResponse{
			ID:    u.ID,
			Email: u.Email,
		})
	}

	return result
}

func GetUserByID(id int) dto.UserResponse {
	user := repository.FindByID(id)

	return dto.UserResponse{
		ID:    user.ID,
		Email: user.Email,
	}
}

func UpdateUser(id int, req dto.UpdateUserRequest) (dto.UserResponse, error) {
	data := repository.FindByID(id)

	if req.Email == "" {
		return dto.UserResponse{}, errors.New("email cannot be empty")
	}

	if len(req.Password) < 8 {
		return dto.UserResponse{}, errors.New("password must be at least 8 characters")
	}
	encoded, err := argon.HashEncoded([]byte(req.Password))

	if err != nil {
		return dto.UserResponse{}, err
	}

	data.Email = req.Email // isi nya diganti
	data.Password = string(encoded)

	repository.Update(data)
	return dto.UserResponse{
		ID:    data.ID,
		Email: data.Email,
	}, nil
}
