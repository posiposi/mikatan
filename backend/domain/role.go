package domain

import (
	"fmt"
	"strings"
)

type Role struct {
	value string
}

func NewRole(value string) (*Role, error) {
	if strings.TrimSpace(value) == "" {
		return nil, fmt.Errorf("role value cannot be empty")
	}

	validRoles := []string{"ADMINISTRATOR", "USER"}
	for _, valid := range validRoles {
		if value == valid {
			return &Role{value: value}, nil
		}
	}

	return nil, fmt.Errorf("invalid role: %s", value)
}

func (r *Role) Value() string {
	return r.value
}

func (r *Role) HasPermission(permission *Permission) bool {
	if permission == nil {
		return false
	}

	switch r.value {
	case "ADMINISTRATOR":
		return true // 管理者は全権限を持つ
	case "USER":
		return permission.Value() == "USER" // ユーザーはユーザー権限のみ
	}

	return false
}

func (r *Role) Equals(other *Role) bool {
	if other == nil {
		return false
	}
	return r.value == other.value
}

func (r *Role) String() string {
	return fmt.Sprintf("Role{Value: %s}", r.value)
}