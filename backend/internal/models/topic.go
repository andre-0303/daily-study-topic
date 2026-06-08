package models

type Topic struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Difficulty  string `json:"difficulty"`
	Description string `json:"description"`
}
