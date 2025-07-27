package handlers

import (
    "encoding/json"
    "net/http"

    "booktracker/db"
    "booktracker/models"
    "booktracker/utils"
)

func Register(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    var input struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }

    err := json.NewDecoder(r.Body).Decode(&input)
    if err != nil || input.Username == "" || input.Password == "" {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    hashedPassword, err := utils.HashPassword(input.Password)
    if err != nil {
        http.Error(w, "Error hashing password", http.StatusInternalServerError)
        return
    }

    var id int
    err = db.DB.QueryRow(
        `INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id`,
        input.Username, hashedPassword,
    ).Scan(&id)

    if err != nil {
        http.Error(w, "Username already exists or DB error", http.StatusConflict)
        return
    }

    json.NewEncoder(w).Encode(models.User{
        ID:       id,
        Username: input.Username,
    })
}
