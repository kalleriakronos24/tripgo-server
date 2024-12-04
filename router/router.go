package router

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	"github.com/kalleriakronos24/khaimal-group/config"
	middleware "github.com/kalleriakronos24/khaimal-group/controllers/middlewares"
	v1 "github.com/kalleriakronos24/khaimal-group/controllers/v1"
	v1Master "github.com/kalleriakronos24/khaimal-group/controllers/v1/master"
	"github.com/kalleriakronos24/khaimal-group/pkg/location"
	"github.com/kalleriakronos24/khaimal-group/utils"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// func corsMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		allowHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization"

// 		w.Header().Set("Content-Type", "application/json")
// 		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4321")
// 		w.Header().Set("Access-Control-Allow-Methods", "POST, PUT, PATCH, GET, DELETE")
// 		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
// 		w.Header().Set("Access-Control-Allow-Credentials", "true")
// 		w.Header().Set("Access-Control-Allow-Headers", allowHeaders)

// 		next.ServeHTTP(w, r)
// 	})
// }

func InitializeRouter() (router *gin.Engine) {
	router = gin.New()
	str := []string{"http://localhost:4321", "http://localhost:3000"}

	if config.AppConfig.Environment == "PRODUCTION" {
		str = []string{"https://wadahgo.com"}
	}

	configCors := cors.DefaultConfig()
	configCors.AddAllowHeaders("Authorization")
	configCors.AllowOrigins = str
	router.Use(cors.New(configCors), middleware.AuthMiddleware)

	v1route := router.Group("/api/v1")
	v1route.Use()
	{

		v1route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

		auth := v1route.Group("/auth")
		{
			auth.POST("/signin", v1.POSTLogin)

		}

		authInternal := v1route.Group("/auth/i")
		{
			authInternal.POST("/signup", v1.POSTLogin)
		}

		authCustomer := v1route.Group("/auth/c")
		{
			authCustomer.POST("/signup", v1.POSTRegisterCustomer)
			authCustomer.POST("/signin", v1.POSTLogin)
		}

		authDriver := v1route.Group("/auth/d")
		{
			authDriver.POST("/signup", v1.POSTRegisterDriver)
			authDriver.POST("/signin", v1.POSTLoginDriver)
		}

		user := v1route.Group("/user")
		{
			user.GET("/", utils.AuthOnly, v1Master.GETAllUser)
			user.GET("/profile", utils.AuthOnly, v1Master.GETUser)
			user.GET("/:id", utils.AuthOnly, v1Master.GETUserByID)

			user.POST("/", utils.AuthOnly, v1.POSTRegisterCustomer)
			user.PUT("/:id", utils.AuthOnly, v1Master.PUTUser)
			user.DELETE("/:id", utils.AuthOnly, v1Master.DELETEUser)
		}

		userDriver := v1route.Group("/user/d")
		{
			userDriver.GET("", utils.AuthOnly, v1Master.GETAllUser)
			userDriver.GET("/profile", utils.AuthOnly, v1Master.GETUser)
			userDriver.GET("/:id", utils.AuthOnly, v1Master.GETUserByID)

			userDriver.POST("/", utils.AuthOnly, v1.POSTRegisterDriver)
			userDriver.PUT("/:id", utils.AuthOnly, v1Master.PUTUser)
			userDriver.DELETE("/:id", utils.AuthOnly, v1Master.DELETEUser)
		}

		bookingTransfer := v1route.Group("/booking/transfer")
		{
			bookingTransfer.GET("/all", utils.AuthOnly, v1.GETAllBookingTransferByCustomer)
			bookingTransfer.POST("/cancel/:id", utils.AuthOnly, v1.POSTCancelBookingTransferByCustomer)
			bookingTransfer.POST("", utils.AuthOnly, v1.POSTBookingTransfer)
		}

		bookingTransferAssigned := v1route.Group("/booking-assigned/transfer")
		{
			bookingTransferAssigned.GET("", utils.AuthOnly, v1.GETAllBookingTransferAssignedByDriverID)
			bookingTransferAssigned.GET("/accepted", utils.AuthOnly, v1.GETAllBookingTransferAccepteddByDriverID)
			bookingTransferAssigned.GET("/cancelled", utils.AuthOnly, v1.GETAllBookingTransferCancelledByDriverID)
			bookingTransferAssigned.GET("/ongoing", utils.AuthOnly, v1.GETAllBookingTransferOngoingByDriverID)
			bookingTransferAssigned.GET("/completed", utils.AuthOnly, v1.GETAllBookingTransferCompletedByDriverID)
			bookingTransferAssigned.POST("/accept/:id", utils.AuthOnly, v1.POSTAcceptBookingTransfer)
			bookingTransferAssigned.POST("/cancel/:id", utils.AuthOnly, v1.POSTCancelBookingTransfer)
			bookingTransferAssigned.POST("/ongoing/:id", utils.AuthOnly, v1.POSTOngoingBookingTransfer)
			bookingTransferAssigned.POST("/pickup/:id", utils.AuthOnly, v1.POSTPickupBookingTransfer)
			bookingTransferAssigned.POST("/complete/:id", utils.AuthOnly, v1.POSTCompleteBookingTransfer)
		}

		bookingTransferRating := v1route.Group("/booking/transfer/rating")
		{
			bookingTransferRating.POST("", utils.AuthOnly, v1.POSTCreateBookingTransferRating)
		}

		carManagement := v1route.Group("/car-management")
		{
			carManagement.GET("/get", utils.AuthOnly, v1.GETAllCarManagementByDriverID)
			carManagement.POST("/create", utils.AuthOnly, v1.POSTCreateCarManagement)
			carManagement.PUT("/update/:id", utils.AuthOnly, v1.PUTUpdateCarManagementByID)
			carManagement.PUT("/set/active/:id", utils.AuthOnly, v1.UPDCarManagementToActiveByID)
			carManagement.PUT("/set/inactive/:id", utils.AuthOnly, v1.UPDCarManagementToInactiveByID)
			carManagement.DELETE("/delete/:id", utils.AuthOnly, v1.DELCarManagementByID)
		}

		driverTopup := v1route.Group("/driver-topup")
		{
			driverTopup.GET("/get", v1.GETAllDriverTopup)
			driverTopup.POST("/create", utils.AuthOnly, v1.POSTCreateDriverTopup)
			driverTopup.GET("/driver/get", utils.AuthOnly, v1.GETAllDriverTopup)
		}

		driverTopupHistory := v1route.Group("/driver-topup-history")
		{
			driverTopupHistory.GET("/get", utils.AuthOnly, v1.GETAllDriverTopupByDriver)
			driverTopupHistory.POST("/approve", v1.POSTApproveDriverTopup)
			driverTopupHistory.POST("/reject", v1.POSTRejectDriverTopup)
		}

		driverWallet := v1route.Group("/driver-balance")
		{
			driverWallet.GET("/get", utils.AuthOnly, v1.GETDWalletInformationByDriver)
		}

		webStatistic := v1route.Group("/web/statistic/customer")
		{
			webStatistic.GET("/booking-count", utils.AuthOnly, v1.GETCountBookingTransferByCustomer)
		}

		websockets := v1route.Group("/ws")
		{
			websockets.GET("/loc/track/:id", v1.LocationTrackingV2)
			websockets.GET("/loc/track/v3/*any", v1.LocationTrackingV3)
			websockets.GET("/loc/track/v2/:id", v1.ListenLocationTracking)
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
	}

	fileServingMainRoute := config.AppConfig.AppUrlStaticFileMainRoute
	//todo improve static file serving security
	workdir, _ := os.Getwd()
	path := filepath.Join(workdir, "../files-uploaded")
	v1route.StaticFS(fileServingMainRoute, http.Dir(path))

	if viper.GetBool("SOCKET_ENABLED") {

		log.Print("socket activated")
		// wt := websocket.Default
		// wt.ReadBufferSize = 4096
		// wt.WriteBufferSize = 4096
		// // wt.HandshakeTimeout = 10
		// wt.CheckOrigin = func(req *http.Request) bool {
		// 	return true
		// }

		// server := socketio.NewServer(&engineio.Options{
		// 	Transports: []transport.Transport{
		// 		&polling.Transport{
		// 			CheckOrigin: func(req *http.Request) bool {
		// 				return true
		// 			},
		// 		},
		// 		&websocket.Transport{
		// 			CheckOrigin: func(req *http.Request) bool {
		// 				return true
		// 			},
		// 		},
		// 	},
		// })
		server := socketio.NewServer(nil)

		if err := server.Serve(); err != nil {
			log.Fatalf("socketio listen error: %s\n", err)
		}

		location.InitChatEndpoints(server)
		go func() {
			if err := server.Serve(); err != nil {
				log.Fatalf("socketio listen error: %s\n", err)
			}
		}()
		defer server.Close()
		router.GET("/socket.io/*any", gin.WrapH(server))
		router.POST("/socket.io/*any", gin.WrapH(server))

		http.Handle("/socket.io/", server)

		log.Fatal(http.ListenAndServe(":5001", nil))
	}
	return router
}
