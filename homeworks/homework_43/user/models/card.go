package models

type Card struct {
	Id     string `json:"id"`
	Number string `json:"number"`
	UserID string `json:"user_id"`
}
