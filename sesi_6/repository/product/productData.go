package product

import (
	"sesi_6/repository/warehouse"
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Id          uint `gorm:"primaryKey"`
	Name        string
	ProductType string
	Stock       uint
	Price       uint64
	WarehouseId uint
	Warehouse   warehouse.Warehouse `gorm:"foreignKey:WarehouseId"`
	CreatedAt   time.Time
	UpdatedAt   *time.Time
	DeletedAt   *time.Time
}
