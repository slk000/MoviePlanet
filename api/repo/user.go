package repo

import (
	"movie_planet/infra"
	"movie_planet/models"
	"movie_planet/util"
)

type UserRepo struct {
	db infra.Database
}

func NewUserRepo(db infra.Database) UserRepo {
	return UserRepo{
		db: db,
	}
}

func (u UserRepo) CreateUser(user models.UserRegister) error {
	var newUser models.User
	newUser.Email = user.Email
	newUser.Name = user.Name
	newUser.Password = user.Password
	return u.db.DB.Create(&newUser).Error
}

func (u UserRepo) LoginUser(user models.UserLogin) (*models.User, error) {
	var dbUser models.User
	email := user.Email
	password := user.Password

	err := u.db.DB.Where("email=? ", email).First(&dbUser).Error
	if err != nil {
		return nil, err
	}

	passwordErr := util.CheckPasswordHash(password, dbUser.Password)
	if passwordErr != nil {
		return nil, passwordErr
	}
	return &dbUser, nil
}
