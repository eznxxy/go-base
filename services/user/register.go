package user

import (
	"github.com/eznxxy/go-base/requests"
	"github.com/eznxxy/go-base/server/builders"

	"golang.org/x/crypto/bcrypt"
)

func (userService *Service) Register(request *requests.RegisterRequest) error {
	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}

	user := builders.NewUserBuilder().SetEmail(request.Email).
		SetName(request.Name).
		SetPassword(string(encryptedPassword)).
		Build()

	return userService.DB.Create(&user).Error
}
