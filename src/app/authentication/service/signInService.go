package service

import (
	"authentication/constants"
	"authentication/repo"
	"authentication/utils"
	"database/sql"
)

func CheckCredentials(username, password string, db *sql.DB) (string, error) {
	utils.LogInfo("Checking credentials for:" + username)
	storedPassword, err := repo.GetStoredPassword(username, db)
	if err != nil {
		utils.LogError("username doesn't exist")
		return "", constants.ErrUserNotFound
	}

	if password != storedPassword {
		utils.LogError("Password mismatch for user: " + username)
		return "", constants.ErrInvalidPass
	}
	utils.LogInfo("login successful for user" + username)
	return "Login successful", nil
}
