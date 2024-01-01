package role

import (
	"github.com/eznxxy/go-base/models"
	"github.com/eznxxy/go-base/requests"
)

func (roleService *Service) Update(role *models.Role, updateRoleRequest *requests.UpdateRoleRequest) {
	role.Name = updateRoleRequest.Name
	role.Permissions = updateRoleRequest.Permissions
	role.HasFullAccess = updateRoleRequest.HasFullAccess
	roleService.DB.Save(role)
}
