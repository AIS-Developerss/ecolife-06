package handlers

import (
	"net/http"
	"ecolife-06/backend/internal/application"
	"github.com/gin-gonic/gin"
)

// TariffHandler обрабатывает HTTP запросы для тарифов
type TariffHandler struct {
	tariffService *application.TariffService
}

// NewTariffHandler создает новый обработчик тарифов
func NewTariffHandler(tariffService *application.TariffService) *TariffHandler {
	return &TariffHandler{tariffService: tariffService}
}

// GetCurrentTariff обрабатывает получение текущего тарифа
func (h *TariffHandler) GetCurrentTariff(c *gin.Context) {
	tariff, err := h.tariffService.GetCurrentTariff(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, tariff)
}

// GetAllTariffs обрабатывает получение всех тарифов
func (h *TariffHandler) GetAllTariffs(c *gin.Context) {
	tariffs, err := h.tariffService.GetAllTariffs(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, tariffs)
}

