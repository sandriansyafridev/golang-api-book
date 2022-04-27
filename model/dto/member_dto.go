package dto

type MemberCreateDTO struct {
	Name string `json:"name" binding:"required"`
}

type MemberUpdateDTO struct {
	ID   int
	Name string `json:"name" binding:"required"`
}
