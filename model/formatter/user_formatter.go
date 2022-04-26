package formatter

import (
	"github.com/sandriansyafridev/golang/api/book/model/entity"
	"github.com/sandriansyafridev/golang/api/book/model/response"
)

func ToUserResponse(user entity.User) (userResponse response.UserResponse) {
	userResponse.ID = user.ID
	userResponse.Name = user.Name
	userResponse.Email = user.Email
	userResponse.Address = user.Address
	userResponse.Telepon = user.Telepon
	userResponse.CreatedAt = user.CreatedAt
	userResponse.UpdatedAt = user.UpdatedAt

	return
}

func ToUsersResponse(users []entity.User) (usersResponse []response.UserResponse) {
	for _, user := range users {
		usersResponse = append(usersResponse, ToUserResponse(user))
	}
	return
}
