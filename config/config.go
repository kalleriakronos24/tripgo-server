package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port        int
	Environment string
	Debug       bool

	DBHost     string
	DBPort     int
	DBDatabase string
	DBUsername string
	DBPassword string
	DBUrl      string

	JWTSecret string

	SOCKETEnabled bool

	MAILHost     string
	MAILPort     int
	MAILUsername string
	MAILPassword string
	MAILFrom     string
}

var AppConfig Config

func InitializeAppConfig() Config {
	viper.SetConfigName(".env") // allow directly reading from .env file
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("/")
	viper.AllowEmptyEnv(true)
	viper.AutomaticEnv()
	_ = viper.ReadInConfig()

	AppConfig.Port = viper.GetInt("PORT")
	AppConfig.Environment = viper.GetString("ENVIRONMENT")
	AppConfig.Debug = viper.GetBool("DEBUG")

	AppConfig.DBHost = viper.GetString("DB_HOST")
	AppConfig.DBPort = viper.GetInt("DB_PORT")
	AppConfig.DBDatabase = viper.GetString("DB_DATABASE")
	AppConfig.DBUsername = viper.GetString("DB_USERNAME")
	AppConfig.DBPassword = viper.GetString("DB_PASSWORD")
	AppConfig.DBUrl = viper.GetString("DATABASE_URL")

	AppConfig.JWTSecret = viper.GetString("JWT_SECRET")
	AppConfig.SOCKETEnabled = viper.GetBool("SOCKET_ENABLED")

	AppConfig.MAILHost = viper.GetString("MAIL_HOST")
	AppConfig.MAILPort = viper.GetInt("MAIL_PORT")
	AppConfig.MAILUsername = viper.GetString("MAIL_PASSWORD")
	AppConfig.MAILFrom = viper.GetString("MAIL_FROM")

	log.Printf("[INIT] configuration loaded")
	return AppConfig
}
