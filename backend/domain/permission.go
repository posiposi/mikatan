package domain

import (
	"fmt"
	"strings"
)

type Permission struct {
	value string
}

func NewPermission(value string) (*Permission, error) {
	if strings.TrimSpace(value) == "" {
		return nil, fmt.Errorf("permission value cannot be empty")
	}

	validPermissions := []string{"ADMIN", "USER"}
	for _, valid := range validPermissions {
		if value == valid {
			return &Permission{value: value}, nil
		}
	}

	return nil, fmt.Errorf("invalid permission: %s", value)
}

func (p *Permission) Value() string {
	return p.value
}

func (p *Permission) Equals(other *Permission) bool {
	if other == nil {
		return false
	}
	return p.value == other.value
}

func (p *Permission) String() string {
	return fmt.Sprintf("Permission{Value: %s}", p.value)
}
