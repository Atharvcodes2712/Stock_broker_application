package service

import (
	"authentication/constants"
	"authentication/models"
	"authentication/repo"
	"authentication/utils"
	"database/sql"
	"errors"
)

func RegisterUser(user models.SignUpRequest, db *sql.DB) (string, error) {
	exists, err := repo.IsUsernameTaken(user.UserName, db)
	if err != nil {
		utils.LogError("error in checking username")
		return "", errors.New("error checking username")
	}
	if exists {
		utils.LogError("username already exists")
		return "", constants.ErrUserAlreadyExists
	}

	err = repo.InsertUser(user, db)
	if err != nil {
		return "", constants.ErrCreatingUser
	}

	return "SignUp successful", nil
}
