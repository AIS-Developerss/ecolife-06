package handlers

import (
	"net/http"
	"ecolife-06/backend/internal/application"
	"github.com/gin-gonic/gin"
)

// ContainerHandler обрабатывает HTTP запросы для контейнеров
type ContainerHandler struct {
	containerService *application.ContainerService
}

// NewContainerHandler создает новый обработчик контейнеров
func NewContainerHandler(containerService *application.ContainerService) *ContainerHandler {
	return &ContainerHandler{containerService: containerService}
}

// GetAllContainers обрабатывает получение всех контейнеров
func (h *ContainerHandler) GetAllContainers(c *gin.Context) {
	containers, err := h.containerService.GetAllContainers(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, containers)
}

// GetContainer обрабатывает получение контейнера по ID
func (h *ContainerHandler) GetContainer(c *gin.Context) {
	id := c.Param("id")
	container, err := h.containerService.GetContainerByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, container)
}

