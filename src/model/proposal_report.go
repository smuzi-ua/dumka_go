package model

import "time"

type ProposalReport struct {
	Id             int       `json:"id" gorm:"default"`
	ProposalId     int       `json:"proposal_id"`
	UserId         int       `json:"user_id"`
	ReportCategory string    `json:"report_category"`
	Comment        string    `json:"comment"`
	Date           time.Time `json:"date" gorm:"default"`
}
