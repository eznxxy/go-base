package seeders

import (
	"strconv"

	"github.com/eznxxy/go-base/models"

	"gorm.io/gorm"
)

type RoleSeeder struct {
	DB *gorm.DB
}

func NewRoleSeeder(db *gorm.DB) *RoleSeeder {
	return &RoleSeeder{DB: db}
}

func (roleSeeder *RoleSeeder) SetRoles() {
	roles := map[int]map[string]string{
		1: {
			"name":            "admin",
			"permissions":     "[]",
			"has_full_access": "1",
		},
		2: {
			"name":            "user",
			"permissions":     "[]",
			"has_full_access": "0",
		},
	}

	for key, value := range roles {
		role := models.Role{}
		roleSeeder.DB.First(&role, key)

		if role.ID == 0 {
			hasFullAccess, err := strconv.ParseBool(value["has_full_access"])
			if err != nil {
				panic(err)
			}
			role.ID = uint(key)
			role.Name = value["name"]
			role.Permissions = value["permissions"]
			role.HasFullAccess = hasFullAccess
			roleSeeder.DB.Create(&role)
		}
	}
}
