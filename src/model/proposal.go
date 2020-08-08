package model

import "time"

type Proposal struct {
	Id        int       `json:"id,omitempty" gorm:"default"`
	UserId    int       `json:"user_id,omitempty"`
	Title     string    `json:"title,omitempty"`
	Text      string    `json:"text,omitempty"`
	Anonymous bool      `json:"anonymous,omitempty"`
	Date      time.Time `json:"date,omitempty" gorm:"default"`
}

func (Proposal) TableName() string {
	return "proposals"
}
