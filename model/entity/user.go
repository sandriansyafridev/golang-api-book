package entity

import "time"

type User struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement;not null"`
	Name      string    `gorm:"type:varchar(255);not null"`
	Email     string    `gorm:"type:varchar(255);not null"`
	Address   string    `gorm:"type:text"`
	Telepon   string    `gorm:"type:varchar(13)"`
	Password  string    `gorm:"type:varchar(255);not null"`
	Token     string    `gorm:"-"`
	CreatedAt time.Time `gorm:"autoCreateTime:milli"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:milli"`
}
