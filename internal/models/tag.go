package models

import (
	"gorm.io/gorm"
)

type Tag struct {
    gorm.Model
    CompanyID uint   `gorm:"index"`
    Name      string `gorm:"type:text;COLLATE NOCASE"`
}