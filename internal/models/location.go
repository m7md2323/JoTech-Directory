package models


type Location struct {
	ID uint `gorm:"primaryKry"`
	CompanyID uint `gorm:"index"`
	City string `gorm:"type:text COLLATE NOCASE"`
	URL string `gorm:"type:text"`

}