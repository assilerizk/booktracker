package models

type Book struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Status   string `json:"status"` // reading, finished, etc.
	UserID   int    `json:"user_id"`
}
