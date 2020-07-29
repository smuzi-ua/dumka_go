package model

import "time"

type ProposalLike struct{
	Id 	       int64  	 `json:"id" gorm:"default"`
	ProposalId int64     `json:"proposal_id"`
	UserId     int64  	 `json:"user_id"` 
	Feedback   string    `json:"feedback"`
	Date       time.Time `json:"date" gorm:"default"`
}