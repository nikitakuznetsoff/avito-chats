package models

type Chat struct {
	ID			uint32		`json:"chat"`
	Name		string		`json:"name"`
	Users		[]uint32	`json:"users"`
	CreatedAt	string		`json:"-"`
}
