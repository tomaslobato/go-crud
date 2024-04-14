package models

type Comment struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
	User    string `json:"user"`
}
