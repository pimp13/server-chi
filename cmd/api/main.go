package main

import (
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/pimp13/server-chi/internal/api"
	"github.com/pimp13/server-chi/pkg/config"
	database "github.com/pimp13/server-chi/pkg/db"
)

func main() {
	db, err := database.NewMySQLStorage(&mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatal("Error of connection db:", err)
	}
	defer db.Close()
	database.DBPing(db)

	server := api.NewServer(config.Envs.ServerPort, db)
	if err := server.Start(); err != nil {
		log.Fatalf("Error running server: %v", err)
	}
}
