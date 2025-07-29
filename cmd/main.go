package main

import (
	"log"
	"net/http"

	"booktracker/config"
	"booktracker/db"
	"booktracker/handlers"
	"booktracker/middleware"
)

func main() {
	config.LoadEnv()
	db.Connect()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Book Tracker API is running."))
	})

	http.HandleFunc("/register", handlers.Register)
    http.HandleFunc("/login", handlers.Login)
    http.HandleFunc("/books", middleware.JWTAuth(handlers.BooksHandler))

	log.Println("🚀 Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
