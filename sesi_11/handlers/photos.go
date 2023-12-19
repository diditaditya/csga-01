package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetPhotos(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.GetString("username")
		c.JSON(http.StatusOK, gin.H{
			"message":  "success",
			"username": username,
		})
	}
}
