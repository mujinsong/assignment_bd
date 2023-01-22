package controller

import (
	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {
	println("hello world")
	c.JSON(200, gin.H{
		"message": "hello world",
	})
}