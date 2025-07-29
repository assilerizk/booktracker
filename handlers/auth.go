package handlers

import (
    "encoding/json"
    "net/http"

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

func Login(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    var input struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }

    if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    user, err := models.GetUserByUsername(input.Username)
    if err != nil {
        http.Error(w, "Invalid username or password", http.StatusUnauthorized)
        return
    }

    if !utils.CheckPasswordHash(input.Password, user.Password) {
        http.Error(w, "Invalid username or password", http.StatusUnauthorized)
        return
    }

    token, err := utils.GenerateJWT(user.ID)
    if err != nil {
        http.Error(w, "Failed to generate token", http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(map[string]string{
        "token": token,
    })
}
