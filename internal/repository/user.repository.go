package repository

import (
	"errors"
	"koda-b6-backend1/internal/model"
)

var users []model.User
var lastID int

func Create(user model.User) model.User {
	lastID++
	user.ID = lastID

	users = append(users, user)

	return user
}

func FindAll() []model.User {
	return users
}

func FindByID(id int) model.User {
	for _, v := range users {
		if v.ID == id {
			return v
		}
	}

	return model.User{}
}

func Update(user model.User) {
	for i, v := range users {
		if v.ID == user.ID {
			users[i] = user
		}
	}
}

func DeleteUserById(id int) error {

	var newUsers []model.User
	found := false

	for _, u := range users {
		if u.ID == id {
			found = true
			continue
		}

		newUsers = append(newUsers, u)
	}

	if !found {
		return errors.New("user not found")
	}

	users = newUsers
	return nil
}
