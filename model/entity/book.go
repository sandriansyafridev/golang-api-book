package entity

import "time"

type Book struct {
	ID          uint64    `gorm:"primaryKey;autoIncrement;not null"`
	Title       string    `gorm:"type:varchar(255);not null"`
	Description string    `gorm:"type:varchar(255);not null"`
	Price       uint64    `gorm:"bigint;not null"`
	Rating      uint64    `gorm:"type:float;not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime:milli"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime:milli"`
}
