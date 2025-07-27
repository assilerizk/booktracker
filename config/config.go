package config

import (
    "log"
    "os"

    "github.com/joho/godotenv"
)

var (
    DBUrl     string
    JwtSecret string
)

func LoadEnv() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    DBUrl = os.Getenv("DB_URL")
    JwtSecret = os.Getenv("JWT_SECRET")

    if DBUrl == "" || JwtSecret == "" {
        log.Fatal("Missing required environment variables")
    }
}
