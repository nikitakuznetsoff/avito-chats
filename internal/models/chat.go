package models

import (
	"time"
)

type Chat struct {
	ID			int			`json:"chat"`
	Name		string		`json:"name"`
	Users		[]int		`json:"users"`
	CreatedAt	time.Time	`json:"created_at"`
}
