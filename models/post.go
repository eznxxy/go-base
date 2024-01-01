package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	UserID  uint
	Title   string `json:"title" gorm:"type:text"`
	Content string `json:"content" gorm:"type:text"`
	User    User   `gorm:"foreignkey:UserID"`
}
