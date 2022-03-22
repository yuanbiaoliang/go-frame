package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yuanbiaoliang/go-frame/service"
)

func main() {
	router := gin.Default()

	service.RegisterHTTPHandle(router)

	if err := router.Run(":8080"); err != nil {
		logrus.WithFields(logrus.Fields{
			"point": "1647944317",
		}).Fatal("router run fail", err.Error())
	}
}
