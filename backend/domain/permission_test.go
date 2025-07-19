package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPermission_WithValidPermission_ShouldReturnPermission(t *testing.T) {
	permission, err := NewPermission("ADMIN")

	assert.NoError(t, err)
	assert.NotNil(t, permission)
	assert.Equal(t, "ADMIN", permission.Value())
}

func TestNewPermission_WithInvalidPermission_ShouldReturnError(t *testing.T) {
	testCases := []struct {
		name       string
		permission string
	}{
		{"empty permission", ""},
		{"whitespace only", "   "},
		{"invalid permission", "INVALID"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			permission, err := NewPermission(tc.permission)

			assert.Error(t, err)
			assert.Nil(t, permission)
		})
	}
}

func TestPermission_Equals_ShouldReturnCorrectResult(t *testing.T) {
	admin1, _ := NewPermission("ADMIN")
	admin2, _ := NewPermission("ADMIN")
	user, _ := NewPermission("USER")

	assert.True(t, admin1.Equals(admin2))
	assert.False(t, admin1.Equals(user))
	assert.False(t, admin1.Equals(nil))
}

func TestPermission_String_ShouldReturnFormattedString(t *testing.T) {
	permission, _ := NewPermission("ADMIN")

	result := permission.String()

	assert.Equal(t, "Permission{Value: ADMIN}", result)
}