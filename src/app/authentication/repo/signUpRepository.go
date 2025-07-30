package repo

import (
	"database/sql"
	"authentication/constants"
	"authentication/models"
	"authentication/utils"
	"authentication/utils/db"
)


func CreateUser(user models.User) error {
	// Checking if username exists
	var exists bool
	err := db.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE username=$1)", user.UserName).Scan(&exists)
	if err != nil {
		utils.LogError("Error checking username: " + err.Error())
		return constants.ErrFetchingUser
	}
	if exists {
		return constants.ErrUserAlreadyExists
	}

	// Checking if email exists
	var emailExists bool
	err = db.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email=$1)", user.Email).Scan(&emailExists)
	if err != nil {
		utils.LogError("Error checking email: " + err.Error())
		return constants.ErrFetchingUser
	}
	if emailExists {
		return constants.ErrEmailAlreadyExists
	}

	// Inserting the user
	_, err = db.DB.Exec(`INSERT INTO users (username, password, email, phone_number, pan_card)
		VALUES ($1, $2, $3, $4, $5)`,
		user.UserName, user.Password, user.Email, user.PhoneNumber, user.PanCard)

	if err != nil {
		utils.LogError("Error inserting user: " + err.Error())
		return constants.ErrCreatingUser
	}

	return nil
}

func ValidateUser(username, password string) (bool, error) {
	var dbPassword string

	err := db.DB.QueryRow("SELECT password FROM users WHERE username=$1", username).Scan(&dbPassword)
	if err == sql.ErrNoRows {
		return false, constants.ErrUserNotFound
	}
	if err != nil {
		utils.LogError("Error validating user: " + err.Error())
		return false, constants.ErrFetchingUser
	}

	if dbPassword != password {
		return false, constants.ErrInvalidCred
	}

	return true, nil
}
