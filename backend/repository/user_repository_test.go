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
	t.Run("should save user token successfully", func(t *testing.T) {
		db := testutil.SetupTestDB(t)
		userRepo := NewUserRepository(db)
		
		userID := uuid.New().String()
		user := &model.User{
			ID:       userID,
			Name:     "Test User",
			Email:    "test-" + userID + "@example.com",
			Password: "hashedpassword",
		}
		err := userRepo.CreateUser(user)
		require.NoError(t, err)

		token := "jwt-token-example"

		err = userRepo.SaveUserToken(user.ID, token)

		assert.NoError(t, err)

		var updatedUser model.User
		err = db.Where("user_id = ?", user.ID).First(&updatedUser).Error
		require.NoError(t, err)
		assert.Equal(t, token, updatedUser.AccessToken)
	})

	t.Run("should return error when user not found", func(t *testing.T) {
		db := testutil.SetupTestDB(t)
		userRepo := NewUserRepository(db)
		
		nonExistentUserID := uuid.New().String()
		token := "jwt-token-example"

		err := userRepo.SaveUserToken(nonExistentUserID, token)

		assert.Error(t, err)
	})
}

func TestUserRepository_DeleteUserToken(t *testing.T) {
	t.Run("should delete user token successfully", func(t *testing.T) {
		db := testutil.SetupTestDB(t)
		userRepo := NewUserRepository(db)
		
		userID := uuid.New().String()
		user := &model.User{
			ID:          userID,
			Name:        "Test User",
			Email:       "test-" + userID + "@example.com",
			Password:    "hashedpassword",
			AccessToken: "existing-token",
		}
		err := userRepo.CreateUser(user)
		require.NoError(t, err)

		err = userRepo.DeleteUserToken(user.ID)

		assert.NoError(t, err)

		var updatedUser model.User
		err = db.Where("user_id = ?", user.ID).First(&updatedUser).Error
		require.NoError(t, err)
		assert.Empty(t, updatedUser.AccessToken)
	})

	t.Run("should return error when user not found", func(t *testing.T) {
		db := testutil.SetupTestDB(t)
		userRepo := NewUserRepository(db)
		
		nonExistentUserID := uuid.New().String()

		err := userRepo.DeleteUserToken(nonExistentUserID)

		assert.Error(t, err)
	})
}
