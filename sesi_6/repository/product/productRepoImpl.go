package product

import (
	"fmt"
	"sesi_6/models"

	"gorm.io/gorm"
)

type ProductRepoImpl struct {
	db *gorm.DB
}

func New(db *gorm.DB) *ProductRepoImpl {
	db.AutoMigrate(&Product{})
	return &ProductRepoImpl{db}
}

func convertDataToModel(raw Product) models.Product {
	return models.Product{
		Id:          raw.Id,
		Name:        raw.Name,
		ProductType: raw.ProductType,
		Stock:       raw.Stock,
		Price:       raw.Price,
		WarehouseId: raw.WarehouseId,
		Warehouse: models.Warehouse{
			Id:      raw.Warehouse.Id,
			Name:    raw.Warehouse.Name,
			Address: raw.Warehouse.Address,
		},
	}
}

func convertModelToData(prod models.Product) Product {
	return Product{
		Id:          prod.Id,
		Name:        prod.Name,
		ProductType: prod.ProductType,
		Stock:       prod.Stock,
		Price:       prod.Price,
		WarehouseId: prod.WarehouseId,
	}
}

func (repo *ProductRepoImpl) FindAll() []models.Product {
	var rawProducts []Product
	repo.db.Preload("Warehouse").Find(&rawProducts)

	products := []models.Product{}
	for _, raw := range rawProducts {
		products = append(products, convertDataToModel(raw))
	}

	return products
}

func (repo *ProductRepoImpl) FindById(id int) *models.Product {
	var rawProduct Product
	result := repo.db.Preload("Warehouse").First(&rawProduct, uint(id))

	if result.Error != nil {
		return nil
	}

	product := convertDataToModel(rawProduct)
	return &product
}

func (repo *ProductRepoImpl) Create(product models.Product) models.Product {
	prodData := convertModelToData(product)
	result := repo.db.Create(&prodData)
	fmt.Println(result)
	created := convertDataToModel(prodData)
	return created
}

func (repo *ProductRepoImpl) Update(product models.Product) (models.Product, error) {
	prodData := convertModelToData(product)

	var found Product
	result := repo.db.First(&found, prodData.Id)
	if result.Error != nil {
		return product, fmt.Errorf("id %v not found", prodData.Id)
	}

	repo.db.Save(&prodData)
	return product, nil
}

func (repo *ProductRepoImpl) Delete(id int) error {
	var found Product
	result := repo.db.First(&found, uint(id))
	if result.Error != nil {
		return fmt.Errorf("id %v not found", id)
	}

	repo.db.Delete(&Product{}, uint(id))
	return nil
}
