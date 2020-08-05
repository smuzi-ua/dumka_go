package model

type ReportCategory struct {
	Id   int    `json:"id" gorm:"default"`
	Name string `json:"name"`
}
