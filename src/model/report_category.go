package model

type ReportCategory struct {
	Id   int    `json:"id" gorm:"default"`
	Name string `json:"name"`
}

func (ReportCategory) TableName() string {
	return "report_categories"
}
