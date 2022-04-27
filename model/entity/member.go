package entity

import "time"

type Member struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement;not null"`
	Name      string `gorm:"type:varchar(255); not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
