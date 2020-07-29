package model

import "time"

type CommentLike struct{
	Id 	       int64  	 `json:"id" gorm:"default"`
	CommentId  int64     `json:"comment_id"`
	UserId     int64  	 `json:"user_id"` 
	Feedback   string    `json:"feedback"`
	Date       time.Time `json:"date" gorm:"default"`
}