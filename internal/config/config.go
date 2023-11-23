package config

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

type Environment string

const (
	LOCAL Environment = "local"
)

func Env() Environment {
	return Environment(viper.GetString("env"))
}

func Name() string {
	return viper.GetString("name")
}

func ServerPort() int {
	return viper.GetInt("server.port")
}

func ServerDebug() bool {
	return viper.GetBool("server.debug")
}

func Address() string {
	return viper.GetString("server.address")
}

func DbDebug() bool {
	return viper.GetBool("db.debug")
}

func DbName() string {
	return viper.GetString("db.name")
}

func DbPassword() string {
	return viper.GetString("db.password")
}

func DbUser() string {
	return viper.GetString("db.user")
}

func DbPort() string {
	return viper.GetString("db.port")
}

func DbHost() string {
	return viper.GetString("db.host")
}

func DbMaxIdleConn() int {
	return viper.GetInt("db.maxIdleConn")
}

func DbMaxOpenConn() int {
	return viper.GetInt("db.maxOpenConn")
}

// ---------- APP

func LocalePath() string {
	return viper.GetString("app.locale.path")
}

func LogLevel() string {
	return viper.GetString("app.log.level")
}

func Init() {
	viper.SetConfigName(getEnv("CONFIG_NAME", "dev"))
	viper.SetConfigType("yaml")    // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./conf")  // optionally look for config in the working directory
	viper.AddConfigPath("../conf") // optionally look for config in the working directory
	err := viper.ReadInConfig()    // Find and read the config file
	if err != nil {                // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
}

func getEnv(key, fallback string) string {
	log.Info("getting environment")
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
