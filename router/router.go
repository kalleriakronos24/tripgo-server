package router

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/odma1/odma-be/config"
	middleware "gitlab.com/odma1/odma-be/controllers/middlewares"
	v1 "gitlab.com/odma1/odma-be/controllers/v1"
	v1Master "gitlab.com/odma1/odma-be/controllers/v1/master"
	"gitlab.com/odma1/odma-be/utils"
	"net/http"
	"os"
	"path/filepath"
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

		operatingActivity := v1route.Group("/operating-activity")
		{
			operatingActivity.GET("/", utils.AuthOnly, v1.GETAllOperatingActivity)
			operatingActivity.GET("/:id", utils.AuthOnly, v1.GETOperatingActivity)

			operatingActivity.POST("/", utils.AuthOnly, v1.POSTOperatingActivity)

			operatingActivity.PUT("/:id", utils.AuthOnly, v1.PUTOperatingActivity)
		}

		operatingActivityProduct := v1route.Group("/operating-activity-product")
		{
			operatingActivityProduct.GET("/", utils.AuthOnly, v1.GETAllOperatingActivityProduct)
			operatingActivityProduct.GET("/:id", utils.AuthOnly, v1.GETOperatingActivityProduct)

			operatingActivityProduct.POST("/", utils.AuthOnly, v1.POSTOperatingActivityProduct)

			operatingActivityProduct.PUT("/:id", utils.AuthOnly, v1.PUTOperatingActivityProduct)
		}

		product := v1route.Group("/product")
		{
			product.GET("/", utils.AuthOnly, v1.GETAllProduct)
			product.GET("/:id", utils.AuthOnly, v1.GETProduct)

			product.POST("/", utils.AuthOnly, v1.POSTProduct)

			product.PUT("/:id", utils.AuthOnly, v1.PUTProduct)
		}

		productHistory := v1route.Group("/product-history")
		{
			productHistory.GET("/", utils.AuthOnly, v1.GETAllProductHistory)
			productHistory.GET("/:id", utils.AuthOnly, v1.GETProductHistory)

			productHistory.POST("/", utils.AuthOnly, v1.POSTProductHistory)

			productHistory.PUT("/:id", utils.AuthOnly, v1.PUTProductHistory)
		}

		quotation := v1route.Group("/quotation")
		{
			quotation.GET("/", utils.AuthOnly, v1.GETAllQuotation)
			quotation.GET("/:id", utils.AuthOnly, v1.GETQuotation)

			quotation.POST("/", utils.AuthOnly, v1.POSTQuotation)

			quotation.PUT("/:id", utils.AuthOnly, v1.PUTQuotation)
		}

		payment := v1route.Group("/payment")
		{
			payment.GET("/", utils.AuthOnly, v1.GETAllPayment)
			payment.GET("/:id", utils.AuthOnly, v1.GETPayment)

			payment.POST("/", utils.AuthOnly, v1.POSTPayment)

			payment.PUT("/:id", utils.AuthOnly, v1.PUTPayment)
		}

		paymentInstallment := v1route.Group("/payment-installment")
		{
			paymentInstallment.GET("/", utils.AuthOnly, v1.GETAllPaymentInstallment)
			paymentInstallment.GET("/:id", utils.AuthOnly, v1.GETPaymentInstallment)

			paymentInstallment.POST("/", utils.AuthOnly, v1.POSTPaymentInstallment)

			paymentInstallment.PUT("/:id", utils.AuthOnly, v1.PUTPaymentInstallment)
		}

		purchaseOrder := v1route.Group("/purchase-order")
		{
			purchaseOrder.GET("/", utils.AuthOnly, v1.GETAllPurchaseOrder)
			purchaseOrder.GET("/:id", utils.AuthOnly, v1.GETPurchaseOrder)

			purchaseOrder.POST("/", utils.AuthOnly, v1.POSTPurchaseOrder)

			purchaseOrder.PUT("/:id", utils.AuthOnly, v1.PUTPurchaseOrder)
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
