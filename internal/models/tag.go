package models



type Tag struct {
	ID uint `gorm:"primaryKey`
	CompanyID uint `gorm:"index`
	Name string `gorm:"type:text COLLATE NOCASE"`
}