package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func LogMiddleware(log *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		log.WithFields(logrus.Fields{
			"method": c.Request.Method,
			"uri":    c.Request.RequestURI,
			"remote": c.ClientIP(),
		}).Info("Request recieved")

		c.Next()

		end := time.Now()
		elapsed := end.Sub(start)
		log.WithFields(logrus.Fields{
			"method":  c.Request.Method,
			"uri":     c.Request.RequestURI,
			"status":  c.Writer.Status(),
			"size":    c.Writer.Size(),
			"elapsed": elapsed,
		}).Info("Response sended")

	}
}
