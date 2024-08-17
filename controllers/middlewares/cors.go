package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/kalleriakronos24/khaimal-group/config"
)

func CORSMiddleware(c *gin.Context) {

	var referer string

	if config.AppConfig.Environment == "DEVELOPMENT" {
		referer = "http://localhost:8000"
	}

	if config.AppConfig.Environment == "PRODUCTION" {
		referer = "http://195.35.20.241"
	}

	c.Header("Access-Control-Allow-Origin", referer)
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH,OPTIONS,GET,PUT,DELETE")

	if c.Request.Method == "OPTIONS" {
		c.Header("Access-Control-Allow-Origin", referer)
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH,OPTIONS,GET,PUT,DELETE")
		c.AbortWithStatus(204)
		return
	}

	c.Next()
}
