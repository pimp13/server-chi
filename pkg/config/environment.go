package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var Envs = initConfig()

type Config struct {
	AppName               string
	PublicHost            string
	ServerPort            string
	DBUser                string
	DBPassword            string
	DBAddress             string
	DBName                string
	JWTExpirationInSecond int64
	JWTKey                string
	Salt                  string
}

func initConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error of loading env file:", err)
	}

	return &Config{
		AppName:               getEnv("APP_NAME", "GO APP"),
		PublicHost:            getEnv("PUBLIC_HOST", "http://localhost"),
		ServerPort:            getEnv("SERVER_PORT", ":8080"),
		DBUser:                getEnv("DB_USER", "username"),
		DBPassword:            getEnv("DB_PASSWORD", "password"),
		DBAddress:             fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "3306")),
		DBName:                getEnv("DB_NAME", "db_name"),
		JWTExpirationInSecond: getEnvAsInt("JWT_EXPIRATION_IN_SECOND", 3600*24*7),
		JWTKey:                getEnv("JWT_KEY", "secret-key-create-secret-key"),
		Salt:                  getEnv("SALT", "salt-test-021"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int64) int64 {
	if value, exists := os.LookupEnv(key); exists {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return defaultValue
		}
		return i
	}
	return defaultValue
}
