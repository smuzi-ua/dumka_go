package model

import "time"

type CommentReport struct {
	Id             int       `json:"id" gorm:"default"`
	CommentId      int       `json:"comment_id"`
	UserId         int       `json:"user_id"`
	ReportCategory string    `json:"report_category"`
	Comment        string    `json:"comment"`
	Date           time.Time `json:"date" gorm:"default"`
}
