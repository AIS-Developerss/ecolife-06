package handlers

import (
	"net/http"
	"ecolife-06/backend/internal/application"
	"github.com/gin-gonic/gin"
)

// ApplicationHandler обрабатывает HTTP запросы для заявок
type ApplicationHandler struct {
	appService *application.ApplicationService
}

// NewApplicationHandler создает новый обработчик заявок
func NewApplicationHandler(appService *application.ApplicationService) *ApplicationHandler {
	return &ApplicationHandler{appService: appService}
}

// CreateApplication обрабатывает создание заявки
// @Summary Создать заявку
// @Description Создает новую заявку на услугу
// @Tags applications
// @Accept json
// @Produce json
// @Param request body application.CreateApplicationRequest true "Данные заявки"
// @Success 201 {object} domain.Application
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/applications [post]
func (h *ApplicationHandler) CreateApplication(c *gin.Context) {
	var req application.CreateApplicationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	app, err := h.appService.CreateApplication(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, app)
}

// GetApplication обрабатывает получение заявки по ID
func (h *ApplicationHandler) GetApplication(c *gin.Context) {
	id := c.Param("id")
	app, err := h.appService.GetApplication(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, app)
}

// GetAllApplications обрабатывает получение всех заявок
func (h *ApplicationHandler) GetAllApplications(c *gin.Context) {
	limit := c.DefaultQuery("limit", "10")
	offset := c.DefaultQuery("offset", "0")

	apps, err := h.appService.GetAllApplications(c.Request.Context(), parseInt(limit), parseInt(offset))
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, apps)
}

// ErrorResponse представляет ответ с ошибкой
type ErrorResponse struct {
	Error string `json:"error"`
}

