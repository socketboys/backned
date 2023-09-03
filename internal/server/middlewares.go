package server

import (
	"github.com/gin-gonic/gin"
)

func RecoveryAPI(c *gin.Context, err interface{}) {
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"msg": err})
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE")

		c.Next()
	}
}
