package service

import (
	"authentication/models"
	"authentication/repo"
)


func RegisterUser(user models.User) error {
	return repo.CreateUser(user)
}

func LoginUser(username, password string) (bool, error) {
	return repo.ValidateUser(username, password)
}
