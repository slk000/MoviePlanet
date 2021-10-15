package service

import (
	"movie_planet/api/repo"
	"movie_planet/models"
	"movie_planet/util"

	"gorm.io/gorm"
)

type UserService struct {
	repo repo.UserRepo
}

func NewUserService(repo repo.UserRepo) UserService {
	return UserService{
		repo: repo,
	}
}

func (u UserService) CreateUser(user models.UserRegister) error {
	user.Password, _ = util.HashPassword(user.Password)
	return u.repo.CreateUser(user)
}

func (u UserService) LoginUser(user models.UserLogin) (*models.User, error) {
	dbUser, err := u.repo.LoginUser(user)
	if err == gorm.ErrRecordNotFound {
		var newUser models.UserRegister
		newUser.Email = user.Email
		newUser.Password = user.Password
		newUser.Name = user.Email
		u.CreateUser(newUser)
		dbUser, err = u.repo.LoginUser(user)
	}
	return dbUser, err
}
