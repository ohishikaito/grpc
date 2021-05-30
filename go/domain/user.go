package domain

import (
	"time"
)

type User struct {
	Id        uint64 `gorm:"primary_key"`
	LastName  string `validate:"required"`
	FirstName string `validate:"required"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
