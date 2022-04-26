package service

import (
	"errors"

	"github.com/sandriansyafridev/golang/api/book/model/dto"
	"github.com/sandriansyafridev/golang/api/book/model/formatter"
	"github.com/sandriansyafridev/golang/api/book/model/mapping"
	"github.com/sandriansyafridev/golang/api/book/model/response"
	"github.com/sandriansyafridev/golang/api/book/repository"
)

type UserService interface {
	FindAll() (usersResponse []response.UserResponse, err error)
	FindByID(UserID int) (userResponse response.UserResponse, err error)
	Delete(UserID int) (err error)
	Create(request dto.UserCreateDTO) (userResponse response.UserResponse, err error)
	Update(request dto.UserUpdateDTO) (userResponse response.UserResponse, err error)
}

type UserServiceImpl struct {
	repository.UserRepository
}

func (userService *UserServiceImpl) Create(request dto.UserCreateDTO) (userResponse response.UserResponse, err error) {

	if isEmailAvailble := userService.UserRepository.IsEmailAvailable(request.Email); !isEmailAvailble {
		return userResponse, errors.New("email is not available")
	} else {
		user := mapping.ToUserCreate(request)
		if userCreated, err := userService.UserRepository.Save(user); err != nil {
			return userResponse, err
		} else {
			userResponse := formatter.ToUserResponse(userCreated)
			return userResponse, nil
		}

	}

}

func (userService *UserServiceImpl) Update(request dto.UserUpdateDTO) (userResponse response.UserResponse, err error) {

	if user, err := userService.UserRepository.FindByID(request.ID); err != nil {
		return userResponse, err
	} else {
		user.Name = request.Name
		user.Email = request.Email
		user.Address = request.Address
		user.Telepon = request.Telepon
		user.Password = request.Password

		if userUpdated, err := userService.UserRepository.Save(user); err != nil {
			return userResponse, err
		} else {
			userResponse := formatter.ToUserResponse(userUpdated)
			return userResponse, nil
		}

	}

}

func (userService *UserServiceImpl) FindAll() (usersResponse []response.UserResponse, err error) {

	if users, err := userService.UserRepository.FindAll(); err != nil {
		return nil, err
	} else {
		usersResponse = formatter.ToUsersResponse(users)
		return usersResponse, nil
	}

}

func (userService *UserServiceImpl) FindByID(UserID int) (userResponse response.UserResponse, err error) {

	if user, err := userService.UserRepository.FindByID(UserID); err != nil {
		return userResponse, err
	} else {
		userResponse = formatter.ToUserResponse(user)
		return userResponse, nil
	}

}
func (userService *UserServiceImpl) Delete(UserID int) (err error) {

	if user, err := userService.UserRepository.FindByID(UserID); err != nil {
		return err
	} else {
		if err := userService.UserRepository.Delete(user); err != nil {
			return err
		} else {
			return nil
		}
	}
}

func NewUserServiceImpl(userRepository repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		UserRepository: userRepository,
	}
}
