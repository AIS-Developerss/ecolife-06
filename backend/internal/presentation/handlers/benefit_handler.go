package handlers

import (
	"net/http"
	"ecolife-06/backend/internal/application"
	"github.com/gin-gonic/gin"
)

// BenefitHandler обрабатывает HTTP запросы для льгот
type BenefitHandler struct {
	benefitService *application.BenefitService
}

// NewBenefitHandler создает новый обработчик льгот
func NewBenefitHandler(benefitService *application.BenefitService) *BenefitHandler {
	return &BenefitHandler{benefitService: benefitService}
}

// GetAllBenefits обрабатывает получение всех льгот
func (h *BenefitHandler) GetAllBenefits(c *gin.Context) {
	benefits, err := h.benefitService.GetAllBenefits(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, benefits)
}

