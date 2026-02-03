package presentation

import (
	"ecolife-06/backend/internal/application"
	"ecolife-06/backend/internal/infrastructure/logger"
	"ecolife-06/backend/internal/presentation/handlers"
	"ecolife-06/backend/internal/presentation/middleware"
	"time"

	"github.com/gin-gonic/gin"
)

// SetupRouter настраивает маршруты
func SetupRouter(
	appService *application.ApplicationService,
	allowedOrigins []string,
	log *logger.Logger,
) *gin.Engine {
	router := gin.Default()

	// Устанавливаем логгер для обработчиков
	handlers.SetLogger(log)

	// Middleware
	router.Use(middleware.CORS(allowedOrigins))
	router.Use(middleware.LoggingMiddleware(log))
	router.Use(gin.Recovery())

	// Handlers
	appHandler := handlers.NewApplicationHandler(appService)

	// API routes
	api := router.Group("/api")
	{
		// Feedback (форма обратной связи) - публичный эндпоинт с rate limiting
		// Ограничение: 10 запросов в минуту с одного IP
		api.POST("/feedback",
			middleware.RateLimitMiddleware(10, 1*time.Minute),
			appHandler.CreateFeedback)
	}

	return router
}

