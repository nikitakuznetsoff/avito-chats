package models

type Message struct {
	ID 			uint32	`json:"-"`
	Chat 		uint32	`json:"chat"`
	Author		uint32	`json:"author"`
	Text 		string	`json:"text"`
	CreatedAt 	string	`json:"-"`
}
