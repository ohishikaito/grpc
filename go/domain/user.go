package domain

import (
	"time"
)

type User struct {
	Uid       string `gorm:"primaryKey"`
	LastName  string `validate:"required"`
	FirstName string `validate:"required"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
