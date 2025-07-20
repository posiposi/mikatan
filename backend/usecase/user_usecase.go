package usecase

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/posiposi/project/backend/domain"
	ormModel "github.com/posiposi/project/backend/internal/orm/model"
	"github.com/posiposi/project/backend/repository"
	"github.com/posiposi/project/backend/usecase/request"
	"golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface {
	SignUp(req request.SignUpRequest) (*domain.User, error)
	Login(req request.LogInRequest) (string, *domain.User, error)
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
	newUser := ormModel.User{
		Id:       id,
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hash),
		Role:     "USER",
		IsAdmin:  false,
	}
	if err := uu.ur.CreateUser(&newUser); err != nil {
		return nil, err
	}
	
	// model.Userからdomain.Userに変換
	userId, err := domain.NewUserId(newUser.Id)
	if err != nil {
		return nil, err
	}
	
	email, err := domain.NewEmail(newUser.Email)
	if err != nil {
		return nil, err
	}
	
	password, err := domain.NewPassword(newUser.Password)
	if err != nil {
		return nil, err
	}
	
	role, err := domain.NewRole(newUser.Role)
	if err != nil {
		return nil, err
	}
	
	domainUser, err := domain.NewUserWithRole(userId, newUser.Name, email, password, role)
	if err != nil {
		return nil, err
	}
	
	return domainUser, nil
}

func (uu *userUsecase) Login(req request.LogInRequest) (string, *domain.User, error) {
	sotredUser := ormModel.User{}
	if err := uu.ur.GetUserByEmail(&sotredUser, req.Email); err != nil {
		return "", nil, err
	}
	err := bcrypt.CompareHashAndPassword([]byte(sotredUser.Password), []byte(req.Password))
	if err != nil {
		return "", nil, err
	}
	
	userId, err := domain.NewUserId(sotredUser.Id)
	if err != nil {
		return "", nil, err
	}
	
	email, err := domain.NewEmail(sotredUser.Email)
	if err != nil {
		return "", nil, err
	}
	
	password, err := domain.NewPassword(sotredUser.Password)
	if err != nil {
		return "", nil, err
	}
	
	role, err := domain.NewRole(sotredUser.Role)
	if err != nil {
		return "", nil, err
	}
	
	domainUser, err := domain.NewUserWithRole(userId, sotredUser.Name, email, password, role)
	if err != nil {
		return "", nil, err
	}
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": sotredUser.Id,
		"exp":     time.Now().Add(time.Hour * 12).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", nil, err
	}
	
	return tokenString, domainUser, nil
}
