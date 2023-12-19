package product

import (
	"fmt"
	"sesi_6/models"
)

// the product entity
// this is where the business logic for product should live
type Products struct {
	data ProductRepo
}

// the factory of product entity
// this acts as custom initialization of the entity
// use this instead of directly initializing the product entity
func New(data ProductRepo) Products {
	return Products{
		data: data,
	}
}

func (prd *Products) FindAll() []models.Product {
	return prd.data.FindAll()
}

func (prd *Products) FindById(id int) *models.Product {
	return prd.data.FindById(id)
}

func (prd *Products) Create(product models.Product) models.Product {
	return prd.data.Create(product)
}

func (prd *Products) Update(id int, product models.Product) (models.Product, error) {
	found := prd.data.FindById(id)
	if found == nil {
		return models.Product{}, fmt.Errorf("id %d not found", id)
	}

	prd.data.Update(product)

	return product, nil
}

func (prd *Products) Delete(id int) {
	prd.data.Delete(id)
}
