package repository

import (
	"github.com/posiposi/project/backend/domain"
	ormModel "github.com/posiposi/project/backend/internal/orm/model"
	"gorm.io/gorm"
)

type IUserRepository interface {
	GetUserByEmail(user *ormModel.User, email string) error
	CreateUser(user *ormModel.User) error
	GetUserByID(userID string) (*domain.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}

func (ur *userRepository) GetUserByEmail(user *ormModel.User, email string) error {
	if err := ur.db.Where("email = ?", email).First(&user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) CreateUser(user *ormModel.User) error {
	if err := ur.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) GetUserByID(userID string) (*domain.User, error) {
	var user ormModel.User
	if err := ur.db.Where("user_id = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}

	userId, err := domain.NewUserId(user.Id)
	if err != nil {
		return nil, err
	}
	
	email, err := domain.NewEmail(user.Email)
	if err != nil {
		return nil, err
	}
	
	password, err := domain.NewPassword(user.Password)
	if err != nil {
		return nil, err
	}
	
	role, err := domain.NewRole(user.Role)
	if err != nil {
		return nil, err
	}

	domainUser, err := domain.NewUserWithRole(userId, user.Name, email, password, role)
	if err != nil {
		return nil, err
	}
	
	return domainUser, nil
}
