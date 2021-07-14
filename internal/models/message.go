package models

type Message struct {
	ID 			int		`json:"id"`
	Chat 		int		`json:"chat"`
	Author		int		`json:"author"`
	Text 		string	`json:"text"`
	CreatedAt 	string	`json:"created_at"`
}
