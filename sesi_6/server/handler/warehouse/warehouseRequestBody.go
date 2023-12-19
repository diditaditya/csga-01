package warehouse

type WarehouseRequest struct {
	Name    *string `json:"name"`
	Address *string `json:"address"`
}
