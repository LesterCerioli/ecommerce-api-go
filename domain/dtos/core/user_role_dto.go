package core

import "github.com/google/uuid"

type UserRoleDTO struct {
	UserID   uuid.UUID `json:"user_id"`
	UserName string    `json:"user_name,omitempty"`
	RoleID   uuid.UUID `json:"role_id"`
	RoleName string    `json:"role_name,omitempty"`
}
