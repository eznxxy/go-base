package repositories

import (
	"github.com/eznxxy/go-base/models"

	"gorm.io/gorm"
)

type RoleRepositoryQ interface {
	GetRoles(roles *[]models.Role)
	GetRole(role *models.Role, id int)
}

type RoleRepository struct {
	DB *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{DB: db}
}

func (roleRepository *RoleRepository) GetRoles(roles *[]models.Role) {
	roleRepository.DB.Find(roles)
}

func (roleRepository *RoleRepository) GetRole(role *models.Role, id int) {
	roleRepository.DB.Where("id = ? ", id).Find(role)
}
