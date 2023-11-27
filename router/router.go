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
			auth.POST("/signin", v1.POSTLogin)
			auth.POST("/signup", utils.AuthOnly, v1.POSTRegister)
			auth.POST("/signup/sa", v1.POSTRegisterSuperAdmin)
		}

		company := v1route.Group("/company")
		{
			company.GET("/", utils.AuthOnly, v1Master.GETAllCompany)
			company.GET("/:id", utils.AuthOnly, v1Master.GETCompany)

			company.POST("/", utils.AuthOnly, v1Master.POSTCompany)

			company.PUT("/:id", utils.AuthOnly, v1Master.PUTCompany)
		}

		client := v1route.Group("/client")
		{
			client.GET("/", utils.AuthOnly, v1Master.GETAllClient)
			client.GET("/:id", utils.AuthOnly, v1Master.GETClient)

			client.POST("/", utils.AuthOnly, v1Master.POSTClient)

			client.PUT("/:id", utils.AuthOnly, v1Master.PUTClient)
		}

		user := v1route.Group("/user")
		{
			user.GET("/:id", utils.AuthOnly, v1Master.GETUser)

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
