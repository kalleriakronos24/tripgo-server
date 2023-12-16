package middleware

import (
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	//cors.New(cors.Config{
	//	AllowOrigins:     []string{"*"},
	//	AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
	//	AllowHeaders:     []string{"Content-Type", "Content-Length", "Authorization", "Origin"},
	//	ExposeHeaders:    []string{"Content-Type", "Content-Length"},
	//	AllowCredentials: true,
	//	AllowWebSockets:  true,
	//	MaxAge:           12 * time.Hour,
	//})
	//c.Next()

	//c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	//c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	//c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	//c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
	//
	//if c.Request.Method == "OPTIONS" {
	//	c.Status(http.StatusOK)
	//	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	//	return
	//}
	//c.Next()

	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
