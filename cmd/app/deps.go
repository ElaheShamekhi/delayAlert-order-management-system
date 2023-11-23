package main

import (
	"delayAlert-order-management-system/db"
	"gorm.io/gorm"
	"log"
)

func postgresDB() *gorm.DB {
	psql, err := db.NewPostgres()
	if err != nil {
		log.Fatalf("failed to initalize db: %v", err)
	}
	return psql
}
