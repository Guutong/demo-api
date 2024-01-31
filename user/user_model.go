package user

import (
	"database/sql"
	"time"
)

type DeletedAt sql.NullTime

type User struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt DeletedAt `gorm:"index"`

	Name string `json:"name" binding:"required"`
}
