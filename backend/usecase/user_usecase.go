package usecase

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/posiposi/project/backend/model"
	"github.com/posiposi/project/backend/repository"
	"golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface {
	SignUp(user model.User) (model.UserResponse, error)
	SignUpWithAutoLogin(user model.User) (model.UserResponse, string, error)
	Login(user model.User) (string, error)
	Logout(userID string) error
}

type userUsecase struct {
	ur repository.IUserRepository
}

func NewUserUsecase(ur repository.IUserRepository) IUserUsecase {
	return &userUsecase{ur}
}

func (uu *userUsecase) SignUp(user model.User) (model.UserResponse, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return model.UserResponse{}, err
	}
	id := uuid.NewString()
	newUser := model.User{ID: id, Email: user.Email, Password: string(hash)}
	if err := uu.ur.CreateUser(&newUser); err != nil {
		return model.UserResponse{}, err
	}
	resUser := model.UserResponse{
		ID:    newUser.ID,
		Email: newUser.Email,
	}
	return resUser, nil
}

func (uu *userUsecase) Login(user model.User) (string, error) {
	sotredUser := model.User{}
	if err := uu.ur.GetUserByEmail(&sotredUser, user.Email); err != nil {
		return "", err
	}
	err := bcrypt.CompareHashAndPassword([]byte(sotredUser.Password), []byte(user.Password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": sotredUser.ID,
		"exp":     time.Now().Add(time.Hour * 12).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}

	if err := uu.ur.SaveUserToken(sotredUser.ID, tokenString); err != nil {
		return "", err
	}

	return tokenString, nil
}

func (uu *userUsecase) SignUpWithAutoLogin(user model.User) (model.UserResponse, string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return model.UserResponse{}, "", err
	}
	id := uuid.NewString()
	newUser := model.User{ID: id, Email: user.Email, Password: string(hash)}
	if err := uu.ur.CreateUser(&newUser); err != nil {
		return model.UserResponse{}, "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": newUser.ID,
		"exp":     time.Now().Add(time.Hour * 12).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return model.UserResponse{}, "", err
	}

	if err := uu.ur.SaveUserToken(newUser.ID, tokenString); err != nil {
		return model.UserResponse{}, "", err
	}

	resUser := model.UserResponse{
		ID:    newUser.ID,
		Email: newUser.Email,
	}
	return resUser, tokenString, nil
}

func (uu *userUsecase) Logout(userID string) error {
	return uu.ur.DeleteUserToken(userID)
}
