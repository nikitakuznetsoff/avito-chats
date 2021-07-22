package models

import (
	"time"
)

type User struct {
	ID 			int			`json:"user"`
	Username 	string		`json:"username"`
	CreatedAt 	time.Time	`json:"created_at"`
}