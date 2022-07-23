package model

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID       uint64 `gorm:"primarykey"`
	CreateAt time.Time
	UpdateAt time.Time
	DelectAt gorm.DeletedAt `gorm:"index"`
}
