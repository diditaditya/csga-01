package warehouse

import (
	"fmt"
	"net/http"
	"reflect"
	"sesi_6/models"
	"sesi_6/warehouse"
	"strconv"

	"github.com/gin-gonic/gin"
)

type WarehouseHandler struct {
	data warehouse.Warehouses
}

func New(wh warehouse.Warehouses) WarehouseHandler {
	return WarehouseHandler{data: wh}
}

// @Summary Find all warehouses
// @Schemes
// @Description Find all warehouses
// @Tags Warehouse
// @Produce json
// @Success 200 {object} []models.Warehouse
// @Router /warehouses [get]
func (handler *WarehouseHandler) FindAll(c *gin.Context) {
	result := handler.data.FindAll()
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    result,
	})
}

// @Summary Find warehouse by id
// @Schemes
// @Description Find warehouse by id
// @Tags Warehouse
// @Produce json
// @Param id path int true "warehouse Id"
// @Success 200 {object} models.Warehouse
// @Router /warehouses/{id} [get]
func (handler *WarehouseHandler) FindById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("id %v is invalid", id),
		})
		return
	}
	result := handler.data.FindById(id)
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    result,
	})
}

// @Summary Create a warehouse
// @Schemes
// @Description Create a warehouse
// @Tags Warehouse
// @Produce json
// @Param warehouse body WarehouseRequest true "Warehouse JSON"
// @Success 201 {object} models.Warehouse
// @Router /warehouses [post]
func (handler *WarehouseHandler) Create(c *gin.Context) {
	var warehouse models.Warehouse
	c.BindJSON(&warehouse)
	result := handler.data.Create(warehouse)
	c.JSON(http.StatusCreated, gin.H{
		"message": "success",
		"data":    result,
	})
}

// @Summary Update a warehouse
// @Schemes
// @Description Update a warehouse
// @Tags Warehouse
// @Produce json
// @Param id path int true "Warehouse Id"
// @Param Warehouse body WarehouseRequest true "Warehouse JSON"
// @Success 201 {object} models.Warehouse
// @Router /warehouses/{id} [put]
func (handler *WarehouseHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("id %v is invalid", id),
		})
		return
	}

	found := handler.data.FindById(id)
	if found == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "warehouse not found",
		})
	}

	var body = WarehouseRequest{}
	err = c.BindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid json request",
		})
		return
	}
	replaceData(body, found)

	result, err := handler.data.Update(id, found)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    result,
	})
}

func replaceData(reqBody WarehouseRequest, target *models.Warehouse) {
	sourceVal := reflect.ValueOf(reqBody)
	targetVal := reflect.ValueOf(target).Elem()

	for i := 0; i < sourceVal.Type().NumField(); i++ {
		fieldName := sourceVal.Field(i).Type().Name()
		val := sourceVal.Field(i)

		if val.IsNil() {
			continue
		}

		switch targetVal.FieldByName(fieldName).Kind() {
		case reflect.String:
			valAddr := val.Interface().(*string)
			targetVal.FieldByName(fieldName).SetString(*valAddr)
		case reflect.Uint:
			valAddr := val.Interface().(*uint)
			val64 := uint64(*valAddr)
			targetVal.FieldByName(fieldName).SetUint(val64)
		case reflect.Uint64:
			valAddr := val.Interface().(*uint64)
			targetVal.FieldByName(fieldName).SetUint(*valAddr)
		}
	}
}

// @Summary Delete warehouse by id
// @Schemes
// @Description Delete warehouse by id
// @Tags Warehouse
// @Produce json
// @Param id path int true "Warehouse Id"
// @Success 200 {object} models.Warehouse
// @Router /warehouses/{id} [delete]
func (handler *WarehouseHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("id %v is invalid", id),
		})
		return
	}
	handler.data.Delete(id)
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
