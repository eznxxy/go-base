package requests

import validation "github.com/go-ozzo/ozzo-validation/v4"

type BasicRole struct {
	Name          string `json:"name" validate:"required" example:"Admin"`
	Permissions   string `json:"permissions" example:"Can:make user"`
	HasFullAccess bool   `json:"has_full_access" example:"1"`
}

func (br BasicRole) Validate() error {
	return validation.ValidateStruct(&br,
		validation.Field(&br.Name, validation.Required),
		validation.Field(&br.Permissions),
		validation.Field(&br.HasFullAccess, validation.Length(1, 1)),
	)
}

type CreateRoleRequest struct {
	BasicRole
}

type UpdateRoleRequest struct {
	BasicRole
}
