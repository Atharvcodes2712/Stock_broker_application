package repo

import (
	"authentication/utils"
	"database/sql"
)

func GetStoredPassword(username string, db *sql.DB) (string, error) {
	var password string
	query := `SELECT password FROM users WHERE username=$1`
	err := db.QueryRow(query, username).Scan(&password)
	if err != nil {
		utils.LogError("user not found in database")
		return "", err
	}
	return password, nil
}
