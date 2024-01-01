package role

import "github.com/eznxxy/go-base/models"

func (roleService *Service) Delete(role *models.Role) {
	roleService.DB.Delete(role)
}
