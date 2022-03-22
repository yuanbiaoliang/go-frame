package service

import (
	"github.com/gin-gonic/gin"
	"github.com/yuanbiaoliang/go-frame/service/test"
)

func RegisterHTTPHandle(router *gin.Engine) {
	t := router.Group("/test")
	{
		t.GET("/get", test.Get)

		t.POST("/post", test.Post)
	}
}
