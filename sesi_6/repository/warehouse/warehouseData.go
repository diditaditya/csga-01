package warehouse

import (
	"time"

	"gorm.io/gorm"
)

type Warehouse struct {
	gorm.Model
	Id        uint `gorm:"primaryKey"`
	Name      string
	Address   string
	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}
