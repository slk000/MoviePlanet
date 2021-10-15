package models

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type User struct {
	ID        int64     `gorm:"primary_key;auto_increment" json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (user *User) TableName() string {
	return "user"
}

type UserLogin struct {
	Email    string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
}
type UserClaims struct {
	User
	jwt.StandardClaims
}

type UserRegister struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Name     string `form:"name" json:"name"`
}

func (user *User) ResponseMap() map[string]interface{} {
	response := make(map[string]interface{})
	response["id"] = user.ID
	response["email"] = user.Email
	response["name"] = user.Name
	response["created_at"] = user.CreatedAt
	response["updated_at"] = user.UpdatedAt

	return response
}
