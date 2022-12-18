package models

//Article table
type Article struct {
	ID         uint     `json:"id" gorm:"primary_key"`
	Title      string   `json:"title"`
	CategoryID uint     `json:"category_id"`
	Category   Category `json:"category"`
}
