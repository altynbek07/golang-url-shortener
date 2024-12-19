package user

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"uniqueIndex"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func NewLink(email string) *User {
	user := &User{
		Email: email,
	}
	return user
}
