package product

import (
	"fmt"
	"net/http"
	"sesi_6/models"
	"sesi_6/product"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	data product.Products
}

func New(prod product.Products) ProductHandler {
	return ProductHandler{data: prod}
}

// @Summary Find all products
// @Schemes
// @Description Find all products
// @Tags Product
// @Produce json
// @Success 200 {object} []models.Product
// @Router /products [get]
func (handler *ProductHandler) FindAll(c *gin.Context) {
	result := handler.data.FindAll()
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    result,
	})
}

// @Summary Find product by id
// @Schemes
// @Description Find product by id
// @Tags Product
// @Produce json
// @Param id path int true "Product Id"
// @Success 200 {object} models.Product
// @Router /products/{id} [get]
func (handler *ProductHandler) FindById(c *gin.Context) {
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

// @Summary Create a product
// @Schemes
// @Description Create a product
// @Tags Product
// @Produce json
// @Param product body ProductRequest true "Product JSON"
// @Success 201 {object} models.Product
// @Router /products [post]
func (handler *ProductHandler) Create(c *gin.Context) {
	var product models.Product
	c.BindJSON(&product)

	created := handler.data.Create(product)
	c.JSON(http.StatusCreated, gin.H{
		"message": "success",
		"data":    created,
	})
}

// @Summary Update a product
// @Schemes
// @Description Update a product
// @Tags Product
// @Produce json
// @Param id path int true "Product Id"
// @Param product body ProductRequest true "Product JSON"
// @Success 201 {object} models.Product
// @Router /products/{id} [put]
func (handler *ProductHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("id %v is invalid", id),
		})
		return
	}

	var body = ProductRequest{}
	err = c.BindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid json request",
		})
		return
	}

	found := handler.data.FindById(id)
	if found == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprintf("id %d not found", id),
		})
		return
	}
	replaceData(body, found)

	result, err := handler.data.Update(id, *found)
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

func replaceData(reqBody ProductRequest, target *models.Product) {
	if reqBody.Name != nil {
		target.Name = *reqBody.Name
	}
	if reqBody.ProductType != nil {
		target.ProductType = *reqBody.ProductType
	}
	if reqBody.Stock != nil {
		target.Stock = *reqBody.Stock
	}
	if reqBody.Price != nil {
		target.Price = *reqBody.Price
	}
	if reqBody.WarehouseId != nil {
		target.WarehouseId = *reqBody.WarehouseId
	}
}

// @Summary Delete product by id
// @Schemes
// @Description Delete product by id
// @Tags Product
// @Produce json
// @Param id path int true "Product Id"
// @Success 200 {object} models.Product
// @Router /products/{id} [delete]
func (handler *ProductHandler) Delete(c *gin.Context) {
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
