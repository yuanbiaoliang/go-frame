package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

//限流中间件
func RateLimit(fillInterval time.Duration, cap, quantum int64) gin.HandlerFunc {
	bucket := ratelimit.NewBucketWithQuantum(fillInterval, cap, quantum)
	return func(c *gin.Context) {
		if bucket.TakeAvailable(1) < 1 {
			logrus.WithFields(logrus.Fields{
				"point":   "1647944318",
				"message": "rate limit, ip : " + c.ClientIP(),
			}).Info()
			c.String(http.StatusForbidden, "rate limit...")
			c.Abort()
			return
		}
		c.Next()
	}
}
