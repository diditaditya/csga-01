package product

type ProductRequest struct {
	Name        *string `json:"name"`
	ProductType *string `json:"type"`
	Stock       *uint   `json:"stock"`
	Price       *uint64 `json:"price"`
	WarehouseId *uint   `json:"warehouseId"`
}
