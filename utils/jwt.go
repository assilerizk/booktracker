package utils

import (
    "time"
    "booktracker/config"
    "github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(userID int) (string, error) {
    claims := jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(24 * time.Hour).Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(config.JwtSecret))
}
