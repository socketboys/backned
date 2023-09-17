package utils

import "github.com/gin-gonic/gin"

func SendError(c *gin.Context, code int, err string) {
	c.AbortWithStatusJSON(code, gin.H{
		"msg": err,
	})

	Logger.Error(err)
}
