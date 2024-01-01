package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name          string `json:"name" gorm:"type:varchar(255);"`
	Permissions   string `json:"permissions" gorm:"type:text;"`
	HasFullAccess bool   `json:"has_full_access" gorm:"type:boolean;"`
}
