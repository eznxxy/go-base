package seeders

import (
	"strconv"

	"github.com/eznxxy/go-base/models"
	"golang.org/x/crypto/bcrypt"

	"gorm.io/gorm"
)

type UserSeeder struct {
	DB *gorm.DB
}

func NewUserSeeder(db *gorm.DB) *UserSeeder {
	return &UserSeeder{DB: db}
}

func (userSeeder *UserSeeder) SetUsers() {
	users := map[int]map[string]string{
		1: {
			"role_id":  "1",
			"email":    "user1@example.com",
			"name":     "user1",
			"password": "password1",
		},
		2: {
			"role_id":  "2",
			"email":    "user2@example.com",
			"name":     "user2",
			"password": "password2",
		},
	}

	for key, value := range users {
		user := models.User{}
		userSeeder.DB.First(&user, key)

		if user.ID == 0 {
			role, err := strconv.Atoi(value["role_id"])
			if err != nil {
				panic(err)
			}

			encryptedPassword, err := bcrypt.GenerateFromPassword(
				[]byte(value["password"]),
				bcrypt.DefaultCost,
			)
			if err != nil {
				panic(err)
			}

			user.ID = uint(key)
			user.RoleID = uint(role)
			user.Email = value["email"]
			user.Name = value["name"]
			user.Password = string(encryptedPassword)
			userSeeder.DB.Create(&user)
		}
	}
}
