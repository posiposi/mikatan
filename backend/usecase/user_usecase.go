package usecase

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/posiposi/project/backend/domain"
	"github.com/posiposi/project/backend/repository"
	"github.com/posiposi/project/backend/usecase/request"
	"golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface {
	SignUp(req request.SignUpRequest) (*domain.User, error)
	Login(req request.LogInRequest) (string, *domain.User, error)
	GetUserById(userId string) (*domain.User, error)
}

type userUsecase struct {
	ur repository.IUserRepository
}

func NewUserUsecase(ur repository.IUserRepository) IUserUsecase {
	return &userUsecase{ur}
}

func (uu *userUsecase) SignUp(req request.SignUpRequest) (*domain.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		return nil, err
	}
	
	id := uuid.NewString()
	userId, err := domain.NewUserId(id)
	if err != nil {
		return nil, err
	}
	
	email, err := domain.NewEmail(req.Email)
	if err != nil {
		return nil, err
	}
	
	password, err := domain.NewPassword(string(hash))
	if err != nil {
		return nil, err
	}
	
	role, err := domain.NewRole("USER")
	if err != nil {
		return nil, err
	}
	
	domainUser, err := domain.NewUserWithRole(userId, req.Name, email, password, role)
	if err != nil {
		return nil, err
	}
	
	if err := uu.ur.CreateUser(domainUser); err != nil {
		return nil, err
	}
	
	return domainUser, nil
}

func (uu *userUsecase) Login(req request.LogInRequest) (string, *domain.User, error) {
	email, err := domain.NewEmail(req.Email)
	if err != nil {
		return "", nil, err
	}
	
	domainUser, err := uu.ur.GetUserByEmail(email)
	if err != nil {
		return "", nil, err
	}
	
	err = bcrypt.CompareHashAndPassword([]byte(domainUser.Password().Value()), []byte(req.Password))
	if err != nil {
		return "", nil, err
	}
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": domainUser.Id().Value(),
		"exp":     time.Now().Add(time.Hour * 12).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", nil, err
	}
	
	return tokenString, domainUser, nil
}

func (uu *userUsecase) GetUserById(userId string) (*domain.User, error) {
	userIdDomain, err := domain.NewUserId(userId)
	if err != nil {
		return nil, err
	}
	
	return uu.ur.GetUserById(userIdDomain)
}
