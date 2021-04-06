package util

import "github.com/gin-gonic/gin"

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Next()
	}

}
