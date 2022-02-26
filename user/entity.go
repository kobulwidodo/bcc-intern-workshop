package user

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Username string `gorm:"unique"`
	Password string
	Bio      string
}

type RegisterInput struct {
	Name     string `binding:"required"`
	Username string `binding:"required"`
	Password string `binding:"required"`
	Bio      string `binding:"required"`
}

type LoginInput struct {
	Username string `binding:"required"`
	Password string `binding:"required"`
}
