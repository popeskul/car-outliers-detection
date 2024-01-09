package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/popeskul/car-outliers-detection/internal/domain"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, ErrorResponse{message})
}

type CheckAgesRequest []domain.Machine

type CheckAgesResponse []domain.Machine
