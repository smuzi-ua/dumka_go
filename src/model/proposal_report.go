package model

import "time"

type ProposalReport struct{
	Id 	       	   int64  	 `json:"id" gorm:"default"`
	ProposalId 	   int64     `json:"proposal_id"`
	UserId     	   int64  	 `json:"user_id"` 
	ReportCategory string    `json:"report_category"`
	Comment	   	   string 	 `json:"comment"`
	Date       	   time.Time `json:"date" gorm:"default"`
}