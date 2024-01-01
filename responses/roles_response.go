package responses

import (
	"github.com/eznxxy/go-base/models"
)

type RoleResponse struct {
	Name          string `json:"name" example:"Admin"`
	Permissions   string `json:"permissions" example:"Can:make user"`
	HasFullAccess bool   `json:"has_full_access" example:"1"`
	ID            uint   `json:"id" example:"1"`
}

func NewRoleResponse(roles []models.Role) *[]RoleResponse {
	roleResponse := make([]RoleResponse, 0)

	for i := range roles {
		roleResponse = append(roleResponse, RoleResponse{
			Name:          roles[i].Name,
			Permissions:   roles[i].Permissions,
			HasFullAccess: roles[i].HasFullAccess,
			ID:            roles[i].ID,
		})
	}

	return &roleResponse
}
