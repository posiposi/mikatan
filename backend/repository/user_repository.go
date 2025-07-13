package repository

import (
	"github.com/posiposi/project/backend/model"
	"gorm.io/gorm"
)

type IUserRepository interface {
	GetUserByEmail(user *model.User, email string) error
	CreateUser(user *model.User) error
	SaveUserToken(userID, token string) error
	DeleteUserToken(userID string) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}

func (ur *userRepository) GetUserByEmail(user *model.User, email string) error {
	if err := ur.db.Where("email = ?", email).First(&user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) CreateUser(user *model.User) error {
	if err := ur.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) SaveUserToken(userID, token string) error {
	result := ur.db.Model(&model.User{}).Where("user_id = ?", userID).Update("access_token", token)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (ur *userRepository) DeleteUserToken(userID string) error {
	result := ur.db.Model(&model.User{}).Where("user_id = ?", userID).Update("access_token", nil)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
