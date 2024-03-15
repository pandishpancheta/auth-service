package db

import (
	"auth-service/pkg/config"
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/lib/pq"
)

func Init(cfg *config.Config) *sql.DB {
	port, err := strconv.Atoi(cfg.DB_PORT)
	if err != nil {
		panic(err)
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable", cfg.DB_HOST,
		port, cfg.DB_USER, cfg.DB_PASS, cfg.DB_NAME)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	return db
}

func InitTables(db *sql.DB) {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id UUID PRIMARY KEY,
		username TEXT,
		password TEXT,
		email TEXT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS contacts (
		id UUID PRIMARY KEY,
		user_id UUID,
		email TEXT,
		phone TEXT,
		instagram TEXT,
		other TEXT,
		FOREIGN KEY (user_id) REFERENCES users(id)
	)`)
	if err != nil {
		panic(err)
	}

	fmt.Println("Tables created successfully")
}
