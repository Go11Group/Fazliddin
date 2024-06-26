package models

type User struct {
	Id    string
	Name  string
	Age   int
	Phone string
	DeletedAt int `json:"deleted_at"`
}