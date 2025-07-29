package handlers

import (
    "encoding/json"
    "net/http"
    "strings"

    "booktracker/middleware"
    "booktracker/models"
)

// BooksHandler handles all /books routes with method switch
func BooksHandler(w http.ResponseWriter, r *http.Request) {
    user, ok := middleware.GetUserFromContext(r.Context())
    if !ok {
        http.Error(w, "Unauthorized: no user in context", http.StatusUnauthorized)
        return
    }

    switch r.Method {
    case http.MethodGet:
        handleGetBooks(w, r, user.ID)
    case http.MethodPost:
        handleAddBook(w, r, user.ID)
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

func handleGetBooks(w http.ResponseWriter, r *http.Request, userID int) {
    books, err := models.GetBooksByUserID(userID)
    if err != nil {
        http.Error(w, "Failed to fetch books", http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(books)
}

func handleAddBook(w http.ResponseWriter, r *http.Request, userID int) {
    var input struct {
        Title  string `json:"title"`
        Author string `json:"author"`
        Status string `json:"status"` // optional: reading, finished, etc.
    }

    err := json.NewDecoder(r.Body).Decode(&input)
    if err != nil || strings.TrimSpace(input.Title) == "" || strings.TrimSpace(input.Author) == "" {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    book := models.Book{
        Title:  input.Title,
        Author: input.Author,
        Status: input.Status,
        UserID: userID,
    }

    id, err := models.AddBook(book)
    if err != nil {
        http.Error(w, "Failed to add book", http.StatusInternalServerError)
        return
    }

    book.ID = id
    json.NewEncoder(w).Encode(book)
}

// Bonus: You can later add GET/PUT/DELETE /books/{id} using Gorilla Mux or manual path parsing
