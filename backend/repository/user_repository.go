package repository

import (
	"github.com/posiposi/project/backend/domain"
	ormModel "github.com/posiposi/project/backend/internal/orm/model"
	"gorm.io/gorm"
)

type IUserRepository interface {
	GetUserByEmail(email *domain.Email) (*domain.User, error)
	CreateUser(user *domain.User) error
	GetUserById(userId *domain.UserId) (*domain.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}

func (ur *userRepository) GetUserByEmail(email *domain.Email) (*domain.User, error) {
	var user ormModel.User
	if err := ur.db.Where("email = ?", email.Value()).First(&user).Error; err != nil {
		return nil, err
	}

	userId, err := domain.NewUserId(user.Id)
	if err != nil {
		return nil, err
	}

	emailDomain, err := domain.NewEmail(user.Email)
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

	domainUser, err := domain.NewUserWithRole(userId, user.Name, emailDomain, password, role)
	if err != nil {
		return nil, err
	}

	return domainUser, nil
}

func (ur *userRepository) CreateUser(user *domain.User) error {
	ormUser := &ormModel.User{
		Id:       user.Id().Value(),
		Name:     user.Name(),
		Email:    user.Email().Value(),
		Password: user.Password().Value(),
		Role:     user.Role().Value(),
	}

	if err := ur.db.Create(ormUser).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) GetUserById(userId *domain.UserId) (*domain.User, error) {
	var user ormModel.User
	if err := ur.db.Where("id = ?", userId.Value()).First(&user).Error; err != nil {
		return nil, err
	}

	userIdDomain, err := domain.NewUserId(user.Id)
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

	domainUser, err := domain.NewUserWithRole(userIdDomain, user.Name, email, password, role)
	if err != nil {
		return nil, err
	}

	return domainUser, nil
}
