package handlers

import (
    "encoding/json"
    "net/http"

    "booktracker/db"
    "booktracker/models"
    "booktracker/utils"
)

// models/user.go
func CreateUser(username, password string) (int, error) {
    var id int
    err := db.DB.QueryRow(
        `INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id`,
        username, password,
    ).Scan(&id)
    return id, err
}

// handlers/auth.go
func Register(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    var input struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }

    if err := json.NewDecoder(r.Body).Decode(&input); err != nil ||
        input.Username == "" || input.Password == "" {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    hashed, err := utils.HashPassword(input.Password)
    if err != nil {
        http.Error(w, "Server error", http.StatusInternalServerError)
        return
    }

    id, err := models.CreateUser(input.Username, hashed)
    if err != nil {
        http.Error(w, "Username exists or DB error", http.StatusConflict)
        return
    }

    json.NewEncoder(w).Encode(models.User{ID: id, Username: input.Username})
}

