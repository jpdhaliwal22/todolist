package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	UserName  string `gorm:"index:idx_name,unique"`
	Password  string
	CreatedAt time.Time // Automatically managed by GORM for creation time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
