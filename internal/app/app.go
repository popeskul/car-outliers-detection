package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/popeskul/car-outliers-detection/internal/config"
	"github.com/popeskul/car-outliers-detection/internal/handlers"
	"github.com/popeskul/car-outliers-detection/internal/server"
	"github.com/popeskul/car-outliers-detection/internal/services"
	"net/http"
	"time"
)

func App(cfg *config.Config) (*server.Server, error) {
	if cfg == nil {
		return nil, fmt.Errorf("config is nil")
	}

	service, err := services.NewService()
	if err != nil {
		return nil, fmt.Errorf("failed to create service: %w", err)
	}

	r := gin.Default()
	h := handlers.NewHandler(service)

	return server.NewServer(&http.Server{
		Addr:           fmt.Sprintf(":%d", cfg.Server.Port),
		Handler:        h.Init(r),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1 MB
	}), nil
}
