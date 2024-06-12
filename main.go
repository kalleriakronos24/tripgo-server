package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/kalleriakronos24/booklap-be/utils"

	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	"github.com/kalleriakronos24/booklap-be/config"
	"github.com/kalleriakronos24/booklap-be/migrations"
	"github.com/kalleriakronos24/booklap-be/pkg/sockets"
	"github.com/kalleriakronos24/booklap-be/router"
	"github.com/kalleriakronos24/booklap-be/services"
	"github.com/spf13/viper"
)

func init() {
	config.InitializeAppConfig()
	if !config.AppConfig.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
}

func runServer() {

	/*
	* set server timezone to Jakarta, Indonesia
	* so is any request from client will converted to our TimeZone
	 */
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		log.Fatalln(err)
		return
	}
	time.Local = loc

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
	if err := s.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}

func main() {
	if viper.GetBool("SOCKET_ENABLED") {
		server := socketio.NewServer(nil)
		var serv = sockets.RunSocketConnection(server)
		http.Handle("/socket.io/", serv)
	}
	runServer()
}
