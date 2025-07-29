package models

import "booktracker/db"

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Status string `json:"status"` // reading, finished, etc.
	UserID int    `json:"user_id"`
}

func GetBooksByUserID(userID int) ([]Book, error) {
	rows, err := db.DB.Query("SELECT id, title, author, status, user_id FROM books WHERE user_id = $1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	books := []Book{}
	for rows.Next() {
		var b Book
		if err := rows.Scan(&b.ID, &b.Title, &b.Author, &b.Status, &b.UserID); err != nil {
			return nil, err
		}
		books = append(books, b)
	}
	return books, nil
}

func AddBook(book Book) (int, error) {
	var id int
	err := db.DB.QueryRow(
		`INSERT INTO books (title, author, status, user_id) VALUES ($1, $2, $3, $4) RETURNING id`,
		book.Title, book.Author, book.Status, book.UserID,
	).Scan(&id)
	return id, err
}
