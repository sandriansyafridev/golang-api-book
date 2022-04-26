package repository

import (
	"errors"

	"github.com/sandriansyafridev/golang/api/book/model/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll() (users []entity.User, err error)
	FindByID(UserID int) (user entity.User, err error)
	Delete(user entity.User) (err error)
	IsEmailAvailable(email string) (available bool)
	Save(user entity.User) (entity.User, error)
}

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func (userRepository *UserRepositoryImpl) Save(user entity.User) (userResult entity.User, err error) {

	if err := userRepository.DB.Save(&user).Error; err != nil {
		return user, errors.New("oke")
	} else {
		return user, nil
	}

}

func (userRepository *UserRepositoryImpl) FindAll() (users []entity.User, err error) {
	if err := userRepository.DB.Find(&users).Error; err != nil {
		return nil, err
	} else {
		return users, nil
	}
}

func (userRepository *UserRepositoryImpl) FindByID(UserID int) (user entity.User, err error) {
	if err := userRepository.DB.First(&user, UserID).Error; err != nil {
		return user, err
	} else {
		return user, nil
	}
}

func (userRepository *UserRepositoryImpl) Delete(user entity.User) (err error) {
	if err := userRepository.DB.Delete(&user).Error; err != nil {
		return err
	} else {
		return nil
	}
}

func (userRepository *UserRepositoryImpl) IsEmailAvailable(email string) (available bool) {

	user := entity.User{}
	if err := userRepository.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return true
	} else {
		return false
	}

}

func NewUserRepositoryImpl(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		DB: db,
	}
}
