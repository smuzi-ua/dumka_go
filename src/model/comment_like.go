package model

import "time"

type CommentLike struct {
	Id        int       `json:"id" gorm:"default"`
	CommentId int       `json:"comment_id"`
	UserId    int       `json:"user_id"`
	Feedback  string    `json:"feedback"`
	Date      time.Time `json:"date" gorm:"default"`
}

func (CommentLike) TableName() string {
	return "comment_likes"
}
