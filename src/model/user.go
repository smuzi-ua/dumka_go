package model

import (
	"github.com/thanhpk/randstr"
	"math/rand"
	"time"
)

// todo add picture
type User struct {
	Id            int        `json:"id,omitempty" gorm:"default"`
	SchoolId      int        `json:"school_id,omitempty"`
	Name          string     `json:"name,omitempty"`
	Nickname      string     `json:"nickname,omitempty"`
	Code          int        `json:"code,omitempty" gorm:"default"`
	CodeGenerated *time.Time `json:"code_generated,omitempty" gorm:"default"`
	Token         string     `json:"token,omitempty" gorm:"default"`
	Type          string     `json:"type,omitempty" gorm:"default"`
	BannedTo      *time.Time `json:"banned_to,omitempty" gorm:"default"`
	Verified      bool       `json:"verified,omitempty" gorm:"default"`
	Date          *time.Time `json:"date,omitempty" gorm:"default"`
}

func (User) TableName() string {
	return "users"
}

func (u *User) UpdateToken() {
	u.Token = randstr.String(32)
}

func (u *User) UpdateCode() {
	u.Code = rand.Intn(1000000) + 100000
	ts := time.Now()
	u.CodeGenerated = &ts
}
