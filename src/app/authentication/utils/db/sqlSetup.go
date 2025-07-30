package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"authentication/utils"
)

var DB *sql.DB

func InitDB() {
	cfg := utils.GetConfig()

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Password, cfg.DB.DBName, cfg.DB.SSLMode)

	var err error
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Error connecting to DB:", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("Error pinging DB:", err)
	}

	log.Println(" Connected to PostgreSQL database")
}
