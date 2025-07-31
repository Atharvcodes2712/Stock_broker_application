package repo

import (
	"authentication/models"
	"authentication/utils"
	"database/sql"
)

func IsUsernameTaken(username string, db *sql.DB) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE username = $1)`
	err := db.QueryRow(query, username).Scan(&exists)
	utils.LogError("username already taken")
	return exists, err
}

func InsertUser(user models.SignUpRequest, db *sql.DB) error {
	query := `
		INSERT INTO users (username, password, email, phone_number, pan_card)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err := db.Exec(query, user.UserName, user.Password, user.Email, user.PhoneNumber, user.PanCard)
	if err != nil {
		return err
	}
	utils.LogInfo("User created")
	return nil
}
