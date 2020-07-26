package model

import "time"

type Comment struct{
	Id 	       int64  	 `json:"id" gorm:"default"`
	ProposalId int64     `json:"proposal_id"`
	UserId     int64  	 `json:"user_id"` 
	Comment	   string 	 `json:"comment"`
	Anonymous  bool    	 `json:"anonymous"`
	Date       time.Time `json:"date" gorm:"default"`
}