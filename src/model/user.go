package model

import (
	"github.com/thanhpk/randstr"
	"math/rand"
	"time"
)

type User struct {
	Id            int       `json:"id" gorm:"default"`
	SchoolId      int       `json:"school_id"`
	Name          string    `json:"name"`
	Nickname      string    `json:"nickname"`
	Code          int       `json:"code" gorm:"default"`
	CodeGenerated time.Time `json:"code_generated" gorm:"default"`
	Token         string    `json:"token" gorm:"default"`
	Type          string    `json:"type" gorm:"default"`
	BannedTo      time.Time `json:"banned_to" gorm:"default"`
	Verified      bool      `json:"verified" gorm:"default"`
	Date          time.Time `json:"date" gorm:"default"`
}

func (u *User) UpdateToken() {
	u.Token = randstr.String(32)
}

func (u *User) UpdateCode() {
	u.Code = rand.Intn(1000000) + 100000
	u.CodeGenerated = time.Now()
}
