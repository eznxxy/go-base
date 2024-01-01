package role

import (
	"github.com/eznxxy/go-base/models"
	"github.com/eznxxy/go-base/requests"

	"gorm.io/gorm"
)

type ServiceWrapper interface {
	Create(role *models.Role)
	Delete(role *models.Role)
	Update(role *models.Role, updateRoleRequest *requests.UpdateRoleRequest)
}

type Service struct {
	DB *gorm.DB
}

func NewRoleService(db *gorm.DB) *Service {
	return &Service{DB: db}
}
