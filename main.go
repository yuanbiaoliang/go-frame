package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yuanbiaoliang/go-frame/middleware"
	"github.com/yuanbiaoliang/go-frame/service"
	"time"
)

func main() {
	router := gin.Default()

	//初始100，每秒放出100
	router.Use(middleware.RateLimit(time.Second, 100, 100))
	router.Use(middleware.CheckParams())

	//注册路由
	service.RegisterHTTPHandle(router)

	if err := router.Run(":8080"); err != nil {
		logrus.WithFields(logrus.Fields{
			"point":   "1647944317",
			"message": "router run fail " + err.Error(),
		}).Fatal()
	}
}
