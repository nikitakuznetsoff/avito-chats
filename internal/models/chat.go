package models

type Chat struct {
	ID			int			`json:"chat"`
	Name		string		`json:"name"`
	Users		[]int		`json:"users"`
	CreatedAt	string		`json:"-"`
}
