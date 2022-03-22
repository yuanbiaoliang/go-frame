package test

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func Get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"type": "get",
	})
}

func Post(c *gin.Context) {

	type postJson struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	var p postJson
	if err := c.BindJSON(&p); err != nil {
		b, _ := c.GetRawData()
		logrus.WithFields(logrus.Fields{
			"point":  "1647949841",
			"err":    "BindJSON err " + err.Error(),
			"params": string(b),
		}).Fatal()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"type": "post",
		"data": p,
	})
}
