package dto

type UserRequestDTO struct {
	Name string `json:"name" form:"name" validate:"required"`
}
