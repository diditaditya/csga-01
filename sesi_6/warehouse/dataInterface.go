package warehouse

import "sesi_6/models"

type WarehouseRepo interface {
	FindAll() []models.Warehouse
	FindById(id int) *models.Warehouse
	Create(warehouse models.Warehouse) models.Warehouse
	Update(warehouse models.Warehouse) (models.Warehouse, error)
	Delete(id int) error
}
