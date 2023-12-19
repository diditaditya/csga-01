package handlers

import (
	"errors"
	"fmt"
	"myGram/auth"
	"myGram/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var creds credentials
		err := c.ShouldBindJSON(&creds)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "bad request",
			})
			return
		}

		var user models.User
		result := db.Where("username = ?", creds.Username).First(&user)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				fmt.Println("user not found!")
				c.JSON(http.StatusUnauthorized, gin.H{
					"error": "invalid username or password",
				})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "something went wrong",
			})
			return
		}

		// if !auth.CheckPassword(creds.Password, user.Password) {
		// 	fmt.Println("incoming password: ", creds.Password)
		// 	fmt.Println("invalid password!")
		// 	c.JSON(http.StatusUnauthorized, gin.H{
		// 		"error": "invalid username or password",
		// 	})
		// 	return
		// }

		token, err := auth.CreateJWT(user.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "something went wrong",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
			"token":   token,
		})
	}
}
