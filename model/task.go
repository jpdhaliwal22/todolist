package model

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	User      User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Detail    string
	Status    string
	CreatedAt time.Time // Automatically managed by GORM for creation time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
