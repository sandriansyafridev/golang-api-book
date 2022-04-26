package dto

type UserUpdateDTO struct {
	ID       int
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Address  string `json:"address" binding:"required"`
	Telepon  string `json:"telepon" binding:"required"`
	Password string `json:"password" binding:"required"`
}
