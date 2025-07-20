package presenter

import (
	"github.com/posiposi/project/backend/domain"
)

type UserResponseJSON struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

type LoginResponseJSON struct {
	Token string           `json:"token"`
	User  UserResponseJSON `json:"user"`
}

type IUserPresenter interface {
	ToJSON(user *domain.User) UserResponseJSON
	ToLoginJSON(token string, user *domain.User) LoginResponseJSON
}

type userPresenter struct{}

func NewUserPresenter() IUserPresenter {
	return &userPresenter{}
}

func (p *userPresenter) ToJSON(user *domain.User) UserResponseJSON {
	return UserResponseJSON{
		Id:    user.Id().Value(),
		Name:  user.Name(),
		Email: user.Email().Value(),
		Role:  user.Role().Value(),
	}
}

func (p *userPresenter) ToLoginJSON(token string, user *domain.User) LoginResponseJSON {
	return LoginResponseJSON{
		Token: token,
		User:  p.ToJSON(user),
	}
}