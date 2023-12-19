package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string `json:"username" validate:"required" gorm:"unique"`
	Email     string `json:"email" validate:"required" gorm:"unique"`
	Password  string `json:"password" validate:"gte=6"`
	Age       int    `json:"age"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Photo struct {
	gorm.Model
	Title    string
	Caption  string
	PhotoUrl string
	UserID   int
	User     User
}

type Comment struct {
	gorm.Model
	UserID    int
	PhotoID   int
	Message   string
	UpdatedAt time.Time
	User      User
	Photo     Photo
}

type SocialMedia struct {
	gorm.Model
	Name           string
	SocialMediaUrl string
	UserID         int
	User           User
}
