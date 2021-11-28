package middlewares

import "github.com/gin-gonic/gin"

func Basic() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Next()
	}
}
