package model

import "time"

type Proposal struct {
	Id        int       `json:"id" gorm:"default"`
	UserId    int       `json:"user_id"`
	Title     string    `json:"title"`
	Text      string    `json:"text"`
	Anonymous bool      `json:"anonymous"`
	Date      time.Time `json:"date" gorm:"default"`
}
