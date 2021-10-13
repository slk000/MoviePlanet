package service

import (
	"movie_planet/api/repo"
	"movie_planet/models"
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
	return u.repo.CreateUser(user)
}

func (u UserService) LoginUser(user models.UserLogin) (*models.User, error) {
	return u.repo.LoginUser(user)
}
