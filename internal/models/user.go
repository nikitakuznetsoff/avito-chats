package models

type User struct {
	ID 			uint32	`json:"user"`
	Username 	string	`json:"username"`
	CreatedAt 	string	`json:"-"`
}