package model

// todo add other models
type School struct {
	Id   int64  `db:"id,pk" json:"id"`
	Name string `db:"name" json:"name"`
}

func (u *School) TableName() string {
	return "schools"
}
