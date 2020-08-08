package model

type School struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Display bool   `json:"display"`
}

func (School) TableName() string {
	return "schools"
}
