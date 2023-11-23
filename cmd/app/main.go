package main

import (
	"delayAlert-order-management-system/internal/config"
	"delayAlert-order-management-system/internal/locale"
	"delayAlert-order-management-system/internal/logger"
	"delayAlert-order-management-system/server"
	"delayAlert-order-management-system/storage"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			postgresDB,

			// Storages
			storage.New,
			// Services
			// Handler

			server.NewServer,
		),

		fx.Supply(),

		fx.Invoke(
			config.Init,
			logger.SetupLogger,
			InitTelemetry,
			locale.Init,
			server.SetupRoutes,
			server.Run,
		),
	)
	app.Run()
}
