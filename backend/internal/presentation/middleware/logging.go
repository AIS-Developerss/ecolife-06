package middleware

import (
	"ecolife-06/backend/internal/infrastructure/logger"
	"time"

	"github.com/gin-gonic/gin"
)

// LoggingMiddleware создает middleware для логирования запросов
func LoggingMiddleware(log *logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method
		ip := c.ClientIP()

		// Обрабатываем запрос
		c.Next()

		// Логируем после обработки
		latency := time.Since(start)
		status := c.Writer.Status()

		fields := map[string]interface{}{
			"method":     method,
			"path":       path,
			"status":     status,
			"latency_ms": latency.Milliseconds(),
			"ip":         ip,
			"user_agent": c.Request.UserAgent(),
		}

		// Добавляем ошибку, если есть
		if len(c.Errors) > 0 {
			fields["error"] = c.Errors.String()
		}

		// Выбираем уровень логирования в зависимости от статуса
		if status >= 500 {
			log.Error("Request failed", fields)
		} else if status >= 400 {
			log.Warn("Request error", fields)
		} else {
			log.Info("Request processed", fields)
		}
	}
}

