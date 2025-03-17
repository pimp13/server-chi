package db

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

func NewMySQLStorage(cfg *mysql.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}

	return db, nil
}

func CheckPing(db *sql.DB) {
	if err := db.Ping(); err != nil {
		log.Fatalf("Error of ping db: %s", err)
		return
	}
	log.Println("Database connection successfully!")
}
