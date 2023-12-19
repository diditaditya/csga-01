package warehouse

import (
	"fmt"
	"sesi_6/models"
)

type Warehouses struct {
	data WarehouseRepo
}

func New(data WarehouseRepo) Warehouses {
	return Warehouses{
		data: data,
	}
}

func (whs *Warehouses) FindAll() []models.Warehouse {
	return whs.data.FindAll()
}

func (whs *Warehouses) FindById(id int) *models.Warehouse {
	return whs.data.FindById(id)
}

func (whs *Warehouses) Create(warehouse models.Warehouse) models.Warehouse {
	return whs.data.Create(warehouse)
}

func (whs *Warehouses) Update(id int, warehouse *models.Warehouse) (models.Warehouse, error) {
	found := whs.FindById(id)
	if found == nil {
		return models.Warehouse{}, fmt.Errorf("id %d not found", id)
	}

	whs.data.Update(*warehouse)

	return *warehouse, nil
}

func (whs *Warehouses) Delete(id int) {
	whs.data.Delete(id)
}
