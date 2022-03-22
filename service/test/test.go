package test

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Get(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"type": "get",
	})
}

func Post(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"type": "post",
	})
}