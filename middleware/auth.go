package middleware

import (
    "context"
    "net/http"
    "strings"
    "booktracker/utils"
    "booktracker/models"
)

type contextKey string

const userContextKey = contextKey("user")

// Middleware function to protect routes by validating JWT token
func JWTAuth(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
            return
        }

        parts := strings.SplitN(authHeader, " ", 2)
        if len(parts) != 2 || parts[0] != "Bearer" {
            http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
            return
        }

        tokenStr := parts[1]

        claims, err := utils.ValidateJWT(tokenStr)
        if err != nil {
            http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
            return
        }

        userIDFloat, ok := claims["user_id"].(float64)
        if !ok {
            http.Error(w, "Invalid token claims", http.StatusUnauthorized)
            return
        }

        userID := int(userIDFloat)

        // Optional: fetch user from DB and add to context (if needed)
        user, err := models.GetUserByID(userID)
        if err != nil {
            http.Error(w, "User not found", http.StatusUnauthorized)
            return
        }

        ctx := context.WithValue(r.Context(), userContextKey, user)
        next(w, r.WithContext(ctx))
    }
}

// Helper to get user from context in handlers
func GetUserFromContext(ctx context.Context) (models.User, bool) {
    user, ok := ctx.Value(userContextKey).(models.User)
    return user, ok
}
