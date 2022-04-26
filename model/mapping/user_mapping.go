package mapping

import (
	"github.com/sandriansyafridev/golang/api/book/model/dto"
	"github.com/sandriansyafridev/golang/api/book/model/entity"
)

func ToUserCreate(request dto.UserCreateDTO) (user entity.User) {
	user.Name = request.Name
	user.Email = request.Email
	user.Address = request.Address
	user.Password = request.Password
	user.Telepon = request.Telepon
	return
}

func ToUserUpdate(request dto.UserUpdateDTO) (user entity.User) {
	user.ID = uint64(request.ID)
	user.Name = request.Name
	user.Email = request.Email
	user.Address = request.Address
	user.Password = request.Password
	user.Telepon = request.Telepon
	return
}
