package model

type Admin struct {
	Id                int    `json:"id" gorm:"default"`
	Telegram_nickname string `json:"telegram_nickname"`
}
