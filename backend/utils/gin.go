package utils

import "github.com/gin-gonic/gin"

func GinErrorMsg(c *gin.Context, err error) {
	c.JSON(500, gin.H{
		"message": err.Error(),
	})
}

func GinErrorMsgTxt(c *gin.Context, err string) {
	c.JSON(500, gin.H{
		"message": err,
	})
}
