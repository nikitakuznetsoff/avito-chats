package models

type User struct {
	ID 			int		`json:"user"`
	Username 	string	`json:"username"`
	CreatedAt 	string	`json:"created_at"`
}