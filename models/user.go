package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	RoleID   uint
	Email    string `json:"email" gorm:"type:varchar(200);"`
	Name     string `json:"name" gorm:"type:varchar(200);"`
	Password string `json:"password" gorm:"type:varchar(200);"`
	Role     Role   `gorm:"foreignkey:RoleID"`
	Post     []Post
}
