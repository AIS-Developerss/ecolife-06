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

// CreateFeedback обрабатывает создание заявки из формы обратной связи
// @Summary Создать заявку из формы обратной связи
// @Description Создает новую заявку с минимальными полями (имя и телефон)
// @Tags feedback
// @Accept json
// @Produce json
// @Param request body application.CreateFeedbackRequest true "Данные формы обратной связи"
// @Success 201 {object} domain.Application
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/feedback [post]
func (h *ApplicationHandler) CreateFeedback(c *gin.Context) {
	var req application.CreateFeedbackRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	app, err := h.appService.CreateFeedback(c.Request.Context(), &req)
	if err != nil {
		status, message := handleError(err)
		c.JSON(status, ErrorResponse{Message: message})
		return
	}

	c.JSON(http.StatusCreated, app)
}

// ErrorResponse представляет ответ с ошибкой
type ErrorResponse struct {
	Message string `json:"message"`
}

