package model

type ReportCategorie struct {
	Id   int64  `json:"id" gorm:"default"`
	Name string `json:"name"`
}
