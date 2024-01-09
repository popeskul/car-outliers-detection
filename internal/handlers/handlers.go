package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/popeskul/car-outliers-detection/docs"
	"github.com/popeskul/car-outliers-detection/internal/services"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handlers struct {
	service services.IService
}

func NewHandler(service services.IService) *Handlers {
	return &Handlers{service: service}
}

func (h *Handlers) Init(router *gin.Engine) *gin.Engine {
	initSwagger(router)

	api := router.Group("/api/")
	{
		handlers := NewHandler(h.service)
		api.POST("/check-ages", handlers.CheckAges)
	}

	return router
}

// @title Your API Title
// @description Your API Description
// @version 1.0
// @host localhost:8080
// @BasePath /api/
func initSwagger(router *gin.Engine) {
	docs.SwaggerInfo.BasePath = "/"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
