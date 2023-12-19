package product

import "sesi_6/models"

type ProductRepo interface {
	FindAll() []models.Product
	FindById(id int) *models.Product
	Create(product models.Product) models.Product
	Update(product models.Product) (models.Product, error)
	Delete(id int) error
}
