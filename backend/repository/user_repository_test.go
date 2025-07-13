package repository

import (
	"testing"

	"github.com/google/uuid"
	"github.com/posiposi/project/backend/model"
	"github.com/posiposi/project/backend/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserRepository_SaveUserToken(t *testing.T) {
	db := testutil.SetupTestDB(t)

	userRepo := NewUserRepository(db)

	t.Run("should save user token successfully", func(t *testing.T) {
		// Arrange
		user := &model.User{
			ID:       uuid.New().String(),
			Name:     "Test User",
			Email:    "test@example.com",
			Password: "hashedpassword",
		}
		err := userRepo.CreateUser(user)
		require.NoError(t, err)

		token := "jwt-token-example"

		// Act
		err = userRepo.SaveUserToken(user.ID, token)

		// Assert
		assert.NoError(t, err)

		// Verify token was saved
		var updatedUser model.User
		err = db.Where("user_id = ?", user.ID).First(&updatedUser).Error
		require.NoError(t, err)
		assert.Equal(t, token, updatedUser.AccessToken)
	})

	t.Run("should return error when user not found", func(t *testing.T) {
		// Arrange
		nonExistentUserID := uuid.New().String()
		token := "jwt-token-example"

		// Act
		err := userRepo.SaveUserToken(nonExistentUserID, token)

		// Assert
		assert.Error(t, err)
	})
}

func TestUserRepository_DeleteUserToken(t *testing.T) {
	db := testutil.SetupTestDB(t)

	userRepo := NewUserRepository(db)

	t.Run("should delete user token successfully", func(t *testing.T) {
		// Arrange
		user := &model.User{
			ID:          uuid.New().String(),
			Name:        "Test User",
			Email:       "test2@example.com",
			Password:    "hashedpassword",
			AccessToken: "existing-token",
		}
		err := userRepo.CreateUser(user)
		require.NoError(t, err)

		// Act
		err = userRepo.DeleteUserToken(user.ID)

		// Assert
		assert.NoError(t, err)

		// Verify token was deleted
		var updatedUser model.User
		err = db.Where("user_id = ?", user.ID).First(&updatedUser).Error
		require.NoError(t, err)
		assert.Empty(t, updatedUser.AccessToken)
	})

	t.Run("should return error when user not found", func(t *testing.T) {
		// Arrange
		nonExistentUserID := uuid.New().String()

		// Act
		err := userRepo.DeleteUserToken(nonExistentUserID)

		// Assert
		assert.Error(t, err)
	})
}