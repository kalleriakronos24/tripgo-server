package main

import (
	"fmt"
	"gitlab.com/odma1/odma-be/utils"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	"github.com/spf13/viper"
	"gitlab.com/odma1/odma-be/config"
	"gitlab.com/odma1/odma-be/migrations"
	"gitlab.com/odma1/odma-be/pkg/sockets"
	"gitlab.com/odma1/odma-be/router"
	"gitlab.com/odma1/odma-be/services"
)

func init() {
	config.InitializeAppConfig()
	if !config.AppConfig.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
}

func runServer() {

	// set application timezone globally
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
	if viper.GetBool("SOCKET_ENABLED") == true {
		server := socketio.NewServer(nil)
		var serv = sockets.RunSocketConnection(server)
		http.Handle("/socket.io/", serv)
	}
	runServer()
}
