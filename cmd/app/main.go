package main

import (
	"delayAlert-order-management-system/handler"
	"delayAlert-order-management-system/internal/config"
	"delayAlert-order-management-system/internal/locale"
	"delayAlert-order-management-system/internal/logger"
	"delayAlert-order-management-system/server"
	"delayAlert-order-management-system/service/delay"
	"delayAlert-order-management-system/storage"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			postgresDB,

			// Clients
			client,

			// Storages
			storage.New,

			// Services
			delay.New,

			// Handlers
			handler.NewDelayHandler,

			server.NewServer,
		),

		fx.Supply(),

		fx.Invoke(
			config.Init,
			logger.SetupLogger,
			InitTelemetry,
			locale.Init,
			server.SetupRoutes,
			handler.SetupDelaysRoutes,
			server.Run,
		),
	)
	app.Run()
}
