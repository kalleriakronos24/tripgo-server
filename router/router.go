package router

import (
	"github.com/gin-gonic/gin"
	middleware "gitlab.com/odma1/odma-be/controllers/middlewares"
	v1 "gitlab.com/odma1/odma-be/controllers/v1"
	v1Master "gitlab.com/odma1/odma-be/controllers/v1/master"
	"gitlab.com/odma1/odma-be/utils"
)

func InitializeRouter() (router *gin.Engine) {
	router = gin.Default()
	v1route := router.Group("/api/v1")
	v1route.Use(
		middleware.CORSMiddleware,
		middleware.AuthMiddleware,
	)
	{
		auth := v1route.Group("/auth")
		{
			auth.POST("/login", v1.POSTLogin).Use()
			auth.POST("/logout", v1.POSTRegister)
		}
		user := v1route.Group("/user")
		{
			user.GET("/:username", utils.AuthOnly, v1Master.GETUser)
			user.PUT("/:id", utils.AuthOnly, v1Master.PUTUser)
		}
		misc := v1route.Group("/misc")
		{
			misc.GET("/ping", v1.Pong)
			misc.POST("/upload", v1.UploadFileSingle)
			misc.POST("/upload-multiple", v1.UploadFileMultiple)
		}
	}
	return
}
