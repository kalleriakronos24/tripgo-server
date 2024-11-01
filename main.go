package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/kalleriakronos24/khaimal-group/migrations"
	"github.com/kalleriakronos24/khaimal-group/utils"

	"github.com/gin-gonic/gin"
	"github.com/kalleriakronos24/khaimal-group/config"
	"github.com/kalleriakronos24/khaimal-group/docs"
	"github.com/kalleriakronos24/khaimal-group/router"
	"github.com/kalleriakronos24/khaimal-group/services"
)

func init() {
	config.InitializeAppConfig()
	if !config.AppConfig.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
}
func runServer() {
	/*
	* set server timezone to KL, Malaysia
	* so is any request from client will converted to our TimeZone
	 */
	loc, err := time.LoadLocation("Asia/Kuala_Lumpur")
	if err != nil {
		log.Fatalln(err)
		return
	}
	time.Local = loc
	// swagger configs
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server Khaimal Group server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:3009"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	//automatic database backup
	if config.AppConfig.Environment == "PRODUCTION" {
		go utils.DatabaseBackupCron()
	}
	// initialize db and migrations
	if err := services.InitializeServices(); err != nil {
		log.Fatalln(err)
	}
	migrations.Migrate()
	// serve all routes and routes configuration
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", config.AppConfig.Port),
		Handler:        router.InitializeRouter(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	//log.Printf("address > %s", s)
	if err := s.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}
func main() {
	runServer()
}
