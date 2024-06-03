package repository

import (
	"a21hc3NpZ25tZW50/model"
	"fmt"

	"gorm.io/gorm"
)

type UserRepository interface {
	Add(user model.User) error
	CheckAvail(user model.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *userRepository {
	return &userRepository{db}
}
func (u *userRepository) Add(user model.User) error {
	// TODO: replace this
	if err := u.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (u *userRepository) CheckAvail(user model.User) error {
	// TODO: replace this
	var count int64
	if err := u.db.Model(&model.User{}).Where("username = ? AND password = ?", user.Username, user.Password).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return nil //=> User Available
	}
	return fmt.Errorf("user not available")
}
