package model

import "time"

type CommentReport struct{
	Id 	       	   int64  	 `json:"id" gorm:"default"`
	CommentId 	   int64     `json:"comment_id"`
	UserId     	   int64  	 `json:"user_id"` 
	ReportCategory string    `json:"report_category"`
	Comment	   	   string 	 `json:"comment"`
	Date           time.Time `json:"date" gorm:"default"`
}