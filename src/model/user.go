package model

import "time"

type User struct {
	Id            int       `json:"id" gorm:"default"`
	SchoolId      int       `json:"school_id"`
	Name          string    `json:"name"`
	Nickname      string    `json:"nickname" gorm:"default"`
	Code          int       `json:"code" gorm:"default"`
	CodeGenerated time.Time `json:"code_generated" gorm:"default"`
	Token         string    `json:"token" gorm:"default"`
	Type          string    `json:"type"`
	BannedTo      time.Time `json:"banned_to" gorm:"default"`
	Verified      bool      `json:"verified" gorm:"default"`
	Date          time.Time `json:"date" gorm:"default"`
}
