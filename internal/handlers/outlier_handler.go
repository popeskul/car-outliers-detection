package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/popeskul/car-outliers-detection/internal/domain"
	"net/http"
)

// CheckAges - detect outliers in machines
// @Summary Detect outliers in machines
// @Description Detect outliers in machines
// @Tags Outlier
// @Accept json
// @Produce json
// @Param machines body CheckAgesRequest true "CheckAgesRequest"
// @Success 200 {object} CheckAgesResponse
// @Failure 400 {object} ErrorResponse
// @Router /api/check-ages [post]
func (h *Handlers) CheckAges(c *gin.Context) {
	var machines []*domain.Machine
	if err := c.BindJSON(&machines); err != nil {
		newErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("error while binding json: %s", err.Error()))
		return
	}

	outliers, err := h.service.OutlierService().DetectOutliers(machines)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("error while detecting outliers: %s", err.Error()))
	}

	outliersValues := make([]domain.Machine, len(outliers))
	for i, outlier := range outliers {
		if outlier != nil {
			outliersValues[i] = *outlier
		}
	}

	c.JSON(http.StatusOK, CheckAgesResponse(outliersValues))

}
