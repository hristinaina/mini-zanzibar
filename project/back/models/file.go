package models

type File struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
	Owner   string `json:"owner"`
}
