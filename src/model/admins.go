package model

type Admin struct {
	Id               int    `json:"id" gorm:"default"`
	TelegramNickname string `json:"telegram_nickname"`
}

func (Admin) TableName() string {
	return "admins"
}
