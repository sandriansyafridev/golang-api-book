package dto

type BookCreateDTO struct {
	Title       string  `json:"title" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Price       uint64  `json:"price" binding:"required"`
	Rating      float64 `json:"rating" binding:"required"`
}
