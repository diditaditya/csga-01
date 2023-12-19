package server

import (
	"net/http"

	"sesi_6/product"
	prodHandler "sesi_6/server/handler/product"
	whHandler "sesi_6/server/handler/warehouse"
	"sesi_6/warehouse"

	_ "sesi_6/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func Start(warehouse warehouse.Warehouses, product product.Products) {

	productHandler := prodHandler.New(product)
	warehouseHandler := whHandler.New(warehouse)

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "up and running",
		})
	})

	productRouter := r.Group("/api/v1/products")
	{
		productRouter.GET("/", productHandler.FindAll)
		productRouter.GET("/:id", productHandler.FindById)
		productRouter.POST("/", productHandler.Create)
		productRouter.PUT("/:id", productHandler.Update)
		productRouter.DELETE("/:id", productHandler.Delete)
	}

	warehouseRouter := r.Group("/api/v1/warehouses")
	{
		warehouseRouter.GET("/", warehouseHandler.FindAll)
		warehouseRouter.GET("/:id", warehouseHandler.FindById)
		warehouseRouter.POST("/", warehouseHandler.Create)
		warehouseRouter.PUT("/:id", warehouseHandler.Update)
		warehouseRouter.DELETE("/:id", warehouseHandler.Delete)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.Run()
}
