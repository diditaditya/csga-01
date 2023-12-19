package handlers

import (
	"myGram/auth"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Authenticated() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")

		jwtToken := strings.Split(header, " ")
		if len(jwtToken) < 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid authorization format",
			})
			return
		}

		if jwtToken[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid authorization scheme",
			})
			return
		}

		username, err := auth.VerifyJWT(jwtToken[1])
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.Set("username", username)
		c.Next()
	}
}
