package router

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/kalleriakronos24/booklap-be/config"
	middleware "github.com/kalleriakronos24/booklap-be/controllers/middlewares"
	v1 "github.com/kalleriakronos24/booklap-be/controllers/v1"
	v1Master "github.com/kalleriakronos24/booklap-be/controllers/v1/master"
	"github.com/kalleriakronos24/booklap-be/utils"
)

func InitializeRouter() (router *gin.Engine) {
	router = gin.Default()

	commonRoute := router.Group("/")

	v1route := router.Group("/api/v1")
	v1route.Use(
		middleware.CORSMiddleware,
		middleware.AuthMiddleware,
	)
	{
		auth := v1route.Group("/auth")
		{
			auth.POST("/signin", v1.POSTLogin)
			auth.POST("/signup/sa", v1.POSTRegisterSuperAdmin)
		}

		// company := v1route.Group("/company")
		// {
		// 	company.GET("/", utils.AuthOnly, v1Master.GETAllCompany)
		// 	company.GET("/:id", utils.AuthOnly, v1Master.GETCompany)
		// 	company.POST("/", utils.AuthOnly, v1Master.POSTCompany)

		// 	company.PUT("/:id", utils.AuthOnly, v1Master.PUTCompany)
		// 	company.DELETE("/:id", utils.AuthOnly, v1Master.DELETECompany)
		// }

		// client := v1route.Group("/client")
		// {
		// 	client.GET("/list", utils.AuthOnly, v1Master.GETAllClients)
		// 	client.GET("/", utils.AuthOnly, v1Master.GETAllClient)
		// 	client.GET("/:id", utils.AuthOnly, v1Master.GETClient)
		// 	client.POST("/", utils.AuthOnly, v1Master.POSTClient)

		// 	client.PUT("/:id", utils.AuthOnly, v1Master.PUTClient)
		// 	client.DELETE("/:id", utils.AuthOnly, v1Master.DELETEClient)
		// }

		user := v1route.Group("/user")
		{
			user.GET("/", utils.AuthOnly, v1Master.GETAllUser)
			user.GET("/profile", utils.AuthOnly, v1Master.GETUser)
			user.GET("/:id", utils.AuthOnly, v1Master.GETUserByID)

			user.POST("/", utils.AuthOnly, v1.POSTRegister)
			user.PUT("/:id", utils.AuthOnly, v1Master.PUTUser)
			user.DELETE("/:id", utils.AuthOnly, v1Master.DELETEUser)
		}

		misc := v1route.Group("/misc")
		{
			misc.GET("/ping", v1.Pong)
			misc.POST("/upload", v1.UploadFileSingle)
			misc.POST("/upload-multiple", v1.UploadFileMultiple)
			if config.AppConfig.Environment == "PRODUCTION" {
				misc.POST("/restore/:fileName", utils.AuthOnly, v1.RestoreDatabase)
			}
		}

		fileServingGroupRoute := config.AppConfig.APPUrlStaticFileGroupRoute
		fileServingMainRoute := config.AppConfig.AppUrlStaticFileMainRoute
		fileServing := commonRoute.Group(fileServingGroupRoute)
		{
			//todo improve static file serving security
			workdir, _ := os.Getwd()
			path := filepath.Join(workdir, "../files-uploaded")
			fileServing.StaticFS(fileServingMainRoute, http.Dir(path))
		}
	}
	return
}
