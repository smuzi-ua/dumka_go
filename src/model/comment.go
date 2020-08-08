package model

import "time"

type Comment struct {
	Id         int       `json:"id" gorm:"default"`
	ProposalId int       `json:"proposal_id"`
	UserId     int       `json:"user_id"`
	Comment    string    `json:"comment"`
	Anonymous  bool      `json:"anonymous"`
	Date       time.Time `json:"date" gorm:"default"`
}

func (Comment) TableName() string {
	return "comments"
}
