package handlers

import (
	"myGram/auth"
	"myGram/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func Register(db *gorm.DB) gin.HandlerFunc {
	validate := validator.New(validator.WithRequiredStructEnabled())

	return func(c *gin.Context) {
		var user models.User

		err := c.ShouldBindJSON(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "bad request",
			})
			return
		}

		err = validate.Struct(user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		password, err := auth.HashPassword(user.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "something went wrong",
			})
			return
		}

		user.Password = password
		db.Save(&user)

		c.JSON(http.StatusCreated, gin.H{
			"message": "success",
			"data": map[string]any{
				"id":       user.ID,
				"username": user.Username,
				"email":    user.Email,
				"age":      user.Age,
			},
		})
	}
}
