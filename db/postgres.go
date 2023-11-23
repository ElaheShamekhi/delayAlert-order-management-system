package db

import (
	"delayAlert-order-management-system/internal/config"
	"delayAlert-order-management-system/storage"
	"fmt"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

var globalDB *gorm.DB

func NewPostgres() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tehran",
		config.DbHost(), config.DbUser(), config.DbPassword(), config.DbName(), config.DbPort(),
	)
	cfg := &gorm.Config{}
	if !config.DbDebug() {
		dbLogger := logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				LogLevel: logger.Info,
				Colorful: true,
			},
		)
		cfg.Logger = dbLogger
	}
	database, err := gorm.Open(postgres.New(postgres.Config{PreferSimpleProtocol: true, DSN: dsn}), cfg)
	if err != nil {
		return nil, fmt.Errorf("postgres database connection failed: %w", err)
	}
	sqlDB, err := database.DB()
	if err != nil {
		return nil, err
	}
	if err = sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping db: %w", err)
	}
	sqlDB.SetMaxIdleConns(config.DbMaxIdleConn())
	sqlDB.SetMaxOpenConns(config.DbMaxOpenConn())
	err = database.AutoMigrate(&storage.Customer{})
	if err != nil {
		return nil, fmt.Errorf("migration failed : %w", err)
	}
	if globalDB == nil {
		globalDB = database
	}
	return database, nil
}
