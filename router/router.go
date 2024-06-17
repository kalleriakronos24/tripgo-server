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
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitializeRouter() (router *gin.Engine) {
	router = gin.Default()

	commonRoute := router.Group("/")

	v1route := router.Group("/api/v1")

	if config.AppConfig.Environment == "DEVELOPMENT" {
		v1route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	v1route.Use(
		middleware.CORSMiddleware,
		middleware.AuthMiddleware,
	)
	{
		auth := v1route.Group("/auth")
		{
			auth.POST("/signin", v1.POSTLogin)
			auth.POST("/signup", v1.POSTRegister)
		}

		authInternal := v1route.Group("/auth/i")
		{
			authInternal.POST("/signin", v1.POSTLogin)
			authInternal.POST("/signup", v1.POSTRegisterInternal)
		}

		authTenant := v1route.Group("/auth/t")
		{
			authTenant.POST("/signin", v1.POSTLogin)
			authTenant.POST("/signup/owner", v1.POSTRegisterBusinessOwner)
			authTenant.POST("/signup/admin", v1.POSTRegisterBusinessAdmin)
		}

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
