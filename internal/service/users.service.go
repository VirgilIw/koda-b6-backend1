package service

import (
	"errors"
	"koda-b6-backend1/internal/model"

	"github.com/matthewhartstonge/argon2"
)

var argon = argon2.DefaultConfig()

var users []model.User // sumber data
var lastID int

func CreateUser(user model.User) (model.User, error) {

	if len(user.Password) < 8 {
		return model.User{}, errors.New("password must be at least 8 characters")
	}

	for _, u := range users {
		if u.Email == user.Email {
			return model.User{}, errors.New("email already registered")
		}
	}

	encoded, err := argon.HashEncoded([]byte(user.Password))
	if err != nil {
		return model.User{}, err
	}

	lastID++
	user.ID = lastID
	user.Password = string(encoded)

	users = append(users, user)

	return user, nil
}

func GetUsers() []model.User {
	return users
}

func GetUserByID(id int) model.User {
	for _, v := range users {
		if v.ID == id {
			return v
		}
	}
	return model.User{}
}
