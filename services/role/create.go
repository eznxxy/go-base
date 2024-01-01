package role

import "github.com/eznxxy/go-base/models"

func (roleService *Service) Create(role *models.Role) {
	roleService.DB.Create(role)
}
