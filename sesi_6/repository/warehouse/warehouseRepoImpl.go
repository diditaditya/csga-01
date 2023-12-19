package warehouse

import (
	"fmt"
	"sesi_6/models"

	"gorm.io/gorm"
)

type WarehouseRepoImpl struct {
	db *gorm.DB
}

func New(db *gorm.DB) *WarehouseRepoImpl {
	db.AutoMigrate(&Warehouse{})
	return &WarehouseRepoImpl{db}
}

func convertDataToModel(raw Warehouse) models.Warehouse {
	return models.Warehouse{Id: raw.Id, Name: raw.Name, Address: raw.Address}
}

func convertModelToData(wh models.Warehouse) Warehouse {
	return Warehouse{Id: wh.Id, Name: wh.Name, Address: wh.Address}
}

func (repo *WarehouseRepoImpl) FindAll() []models.Warehouse {
	var rawWarehouses []Warehouse
	repo.db.Find(&rawWarehouses)

	warehouses := []models.Warehouse{}
	for _, raw := range rawWarehouses {
		warehouses = append(warehouses, convertDataToModel(raw))
	}
	return warehouses
}

func (repo *WarehouseRepoImpl) FindById(id int) *models.Warehouse {
	var rawWarehouse Warehouse
	result := repo.db.First(&rawWarehouse, uint(id))

	if result.Error != nil {
		return nil
	}

	wh := convertDataToModel(rawWarehouse)
	return &wh
}

func (repo *WarehouseRepoImpl) Create(warehouse models.Warehouse) models.Warehouse {
	raw := convertModelToData(warehouse)
	repo.db.Create(&raw)
	wh := convertDataToModel(raw)
	return wh
}

func (repo *WarehouseRepoImpl) Update(warehouse models.Warehouse) (models.Warehouse, error) {
	raw := convertModelToData(warehouse)

	var found Warehouse
	result := repo.db.First(&found, raw.Id)
	if result.Error != nil {
		return warehouse, fmt.Errorf("id %v not found", raw.Id)
	}

	repo.db.Save(&raw)
	return warehouse, nil
}

func (repo *WarehouseRepoImpl) Delete(id int) error {
	var found Warehouse
	result := repo.db.First(&found, uint(id))
	if result.Error != nil {
		return fmt.Errorf("id %v not found", id)
	}

	repo.db.Delete(&Warehouse{}, uint(id))
	return nil
}
