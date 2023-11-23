package server

import (
	"context"
	"delayAlert-order-management-system/internal/config"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/fx"
	"net/http"
)

type Server struct {
	Engine     *gin.Engine
	healthFunc func(ctx *gin.Context)
}

func NewServer() *Server {
	if !config.ServerDebug() {
		gin.SetMode(gin.ReleaseMode)
	}
	s := &Server{Engine: gin.Default(), healthFunc: Health}
	s.Engine.Use(WithTraceID())
	s.Engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	s.setDoc()
	return s
}

func (s *Server) SetHealthFunc(f func() error) *Server {
	s.healthFunc = func(ctx *gin.Context) {
		if err := f(); err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
	}
	return s
}

func SetupRoutes(s *Server) {
	s.Engine.GET("/health", s.healthFunc)
}

func (s *Server) setDoc() {
	//docs.SwaggerInfo.Title = "delayAlert-order-management-system"
	////docs.SwaggerInfo.Description = getDocDescription()
	//docs.SwaggerInfo.Version = "1.0"
	//docs.SwaggerInfo.Host = config.ServerAddress()
}

func Run(lc fx.Lifecycle, s *Server) {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.ServerPort()),
		Handler: s.Engine,
	}
	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			log.Info().Msg("shutting down the server ...")
			return srv.Shutdown(ctx)
		},
		OnStart: func(ctx context.Context) error {
			log.Info().Msg("running server ...")
			go func() {
				if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
					log.Error().Err(err).Msg("failed to run web server")
				}
			}()
			return nil
		}},
	)
}
