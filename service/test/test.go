package test

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

func Get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"type": "get",
	})
}

func Post(c *gin.Context) {
	//Gin框架，body参数只能读取一次
	b, _ := c.GetRawData()
	jsonStr := string(b)
	logrus.WithFields(logrus.Fields{"point": "1647949841", "params": jsonStr}).Info()

	//把读过的字节流重新放到body
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(b))

	//参数绑定
	type postJson struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	var p postJson
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1003,
			"type": "post",
			"msg":  err.Error(),
			"data": jsonStr,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"type": "post",
		"data": p,
	})
}
