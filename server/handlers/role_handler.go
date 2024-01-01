package handlers

import (
	"net/http"
	"strconv"

	"github.com/eznxxy/go-base/models"
	"github.com/eznxxy/go-base/repositories"
	"github.com/eznxxy/go-base/requests"
	"github.com/eznxxy/go-base/responses"
	s "github.com/eznxxy/go-base/server"
	roleservice "github.com/eznxxy/go-base/services/role"

	"github.com/labstack/echo/v4"
)

type RoleHandlers struct {
	server *s.Server
}

func NewRoleHandlers(server *s.Server) *RoleHandlers {
	return &RoleHandlers{server: server}
}

// CreateRole godoc
//
//	@Summary		Create role
//	@Description	Create role
//	@ID				roles-create
//	@Tags			Roles Actions
//	@Accept			json
//	@Produce		json
//	@Param			params	body		requests.CreateRoleRequest	true	"Role title and content"
//	@Success		201		{object}	responses.Data
//	@Failure		400		{object}	responses.Error
//	@Security		ApiKeyAuth
//	@Router			/roles [post]
func (p *RoleHandlers) CreateRole(c echo.Context) error {
	createRoleRequest := new(requests.CreateRoleRequest)

	if err := c.Bind(createRoleRequest); err != nil {
		return err
	}

	if err := createRoleRequest.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Required fields are empty")
	}

	role := models.Role{
		Name:          createRoleRequest.Name,
		Permissions:   createRoleRequest.Permissions,
		HasFullAccess: createRoleRequest.HasFullAccess,
	}
	roleService := roleservice.NewRoleService(p.server.DB)
	roleService.Create(&role)

	return responses.MessageResponse(c, http.StatusCreated, "Role successfully created")
}

// DeleteRole godoc
//
//	@Summary		Delete role
//	@Description	Delete role
//	@ID				roles-delete
//	@Tags			Roles Actions
//	@Param			id	path		int	true	"Role ID"
//	@Success		204	{object}	responses.Data
//	@Failure		404	{object}	responses.Error
//	@Security		ApiKeyAuth
//	@Router			/roles/{id} [delete]
func (p *RoleHandlers) DeleteRole(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	role := models.Role{}

	roleRepository := repositories.NewRoleRepository(p.server.DB)
	roleRepository.GetRole(&role, id)

	if role.ID == 0 {
		return responses.ErrorResponse(c, http.StatusNotFound, "Role not found")
	}

	roleService := roleservice.NewRoleService(p.server.DB)
	roleService.Delete(&role)

	return responses.MessageResponse(c, http.StatusOK, "Role deleted successfully")
}

// GetRoles godoc
//
//	@Summary		Get roles
//	@Description	Get the list of all roles
//	@ID				roles-get
//	@Tags			Roles Actions
//	@Produce		json
//	@Success		200	{array}	responses.RoleResponse
//	@Security		ApiKeyAuth
//	@Router			/roles [get]
func (p *RoleHandlers) GetRoles(c echo.Context) error {
	var roles []models.Role

	roleRepository := repositories.NewRoleRepository(p.server.DB)
	roleRepository.GetRoles(&roles)

	response := responses.NewRoleResponse(roles)
	return responses.Response(c, http.StatusOK, response)
}

// UpdateRole godoc
//
//	@Summary		Update role
//	@Description	Update role
//	@ID				roles-update
//	@Tags			Roles Actions
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int							true	"Role ID"
//	@Param			params	body		requests.UpdateRoleRequest	true	"Role title and content"
//	@Success		200		{object}	responses.Data
//	@Failure		400		{object}	responses.Error
//	@Failure		404		{object}	responses.Error
//	@Security		ApiKeyAuth
//	@Router			/roles/{id} [put]
func (p *RoleHandlers) UpdateRole(c echo.Context) error {
	updateRoleRequest := new(requests.UpdateRoleRequest)
	id, _ := strconv.Atoi(c.Param("id"))

	if err := c.Bind(updateRoleRequest); err != nil {
		return err
	}

	if err := updateRoleRequest.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Required fields are empty")
	}

	role := models.Role{}

	roleRepository := repositories.NewRoleRepository(p.server.DB)
	roleRepository.GetRole(&role, id)

	if role.ID == 0 {
		return responses.ErrorResponse(c, http.StatusNotFound, "Role not found")
	}

	roleService := roleservice.NewRoleService(p.server.DB)
	roleService.Update(&role, updateRoleRequest)

	return responses.MessageResponse(c, http.StatusOK, "Role successfully updated")
}
