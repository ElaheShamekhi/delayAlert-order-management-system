package main

import (
	"context"
	"delayAlert-order-management-system/internal/config"
	"errors"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
	"go.uber.org/fx"
	"os"
)

func InitTelemetry(lc fx.Lifecycle) {
	if config.Env() == config.LOCAL {
		return
	}
	var tp *trace.TracerProvider
	var mp *metric.MeterProvider
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			r := resource.Default()
			r, err := resource.Merge(
				r,
				resource.NewWithAttributes(
					r.SchemaURL(),
					semconv.DeploymentEnvironment(string(config.Env())),
					semconv.ServiceName(config.Name()),
					semconv.ServiceInstanceID(os.Getenv("MY_NODE_NAME")),
					semconv.K8SNodeName(os.Getenv("MY_NODE_NAME")),
					semconv.K8SPodName(os.Getenv("MY_POD_NAME")),
					semconv.K8SNamespaceName(os.Getenv("MY_POD_NAMESPACE")),
					semconv.ServiceNamespace(os.Getenv("MY_POD_NAMESPACE")),
					semconv.ServiceVersion("v1.0")),
			)
			if err != nil {
				log.Error().Err(err).Msg("failed to merge otel resources")
				return nil
			}
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return errors.Join(tp.Shutdown(ctx), mp.Shutdown(ctx))
		},
	})
}
