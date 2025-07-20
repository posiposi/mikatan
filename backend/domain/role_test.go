package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRole_WithValidRole_ShouldReturnRole(t *testing.T) {
	role, err := NewRole("ADMINISTRATOR")

	assert.NoError(t, err)
	assert.NotNil(t, role)
	assert.Equal(t, "ADMINISTRATOR", role.Value())
}

func TestNewRole_WithInvalidRole_ShouldReturnError(t *testing.T) {
	testCases := []struct {
		name string
		role string
	}{
		{"empty role", ""},
		{"whitespace only", "   "},
		{"invalid role", "INVALID_ROLE"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			role, err := NewRole(tc.role)

			assert.Error(t, err)
			assert.Nil(t, role)
		})
	}
}

func TestRole_HasPermission_ShouldReturnCorrectResult(t *testing.T) {
	adminRole, _ := NewRole("ADMINISTRATOR")
	userRole, _ := NewRole("USER")
	adminPermission, _ := NewPermission("ADMIN")
	userPermission, _ := NewPermission("USER")

	assert.True(t, adminRole.HasPermission(adminPermission))
	assert.True(t, adminRole.HasPermission(userPermission)) // 管理者はユーザー権限も持つ
	assert.False(t, userRole.HasPermission(adminPermission))
	assert.True(t, userRole.HasPermission(userPermission))
}

func TestRole_Equals_ShouldReturnCorrectResult(t *testing.T) {
	admin1, _ := NewRole("ADMINISTRATOR")
	admin2, _ := NewRole("ADMINISTRATOR")
	user, _ := NewRole("USER")

	assert.True(t, admin1.Equals(admin2))
	assert.False(t, admin1.Equals(user))
	assert.False(t, admin1.Equals(nil))
}

func TestRole_String_ShouldReturnFormattedString(t *testing.T) {
	role, _ := NewRole("ADMINISTRATOR")

	result := role.String()

	assert.Equal(t, "Role{Value: ADMINISTRATOR}", result)
}