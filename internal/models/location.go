package models

import (
	"gorm.io/gorm"
)

type Location struct {
    gorm.Model
    CompanyID uint   `gorm:"index"`
    City      string `gorm:"type:text;COLLATE NOCASE"`
    URL       string `gorm:"type:text"`
}
