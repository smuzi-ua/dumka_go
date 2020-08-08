package model

import "time"

type ProposalLike struct {
	Id         int       `json:"id" gorm:"default"`
	ProposalId int       `json:"proposal_id"`
	UserId     int       `json:"user_id"`
	Feedback   string    `json:"feedback"`
	Date       time.Time `json:"date" gorm:"default"`
}

func (ProposalLike) TableName() string {
	return "proposal_likes"
}
