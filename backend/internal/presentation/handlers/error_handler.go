package handlers

import (
	"database/sql"
	"ecolife-06/backend/internal/infrastructure/logger"
	"errors"
	"net/http"
	"strings"
)

var appLogger *logger.Logger

// SetLogger устанавливает логгер для обработчиков
func SetLogger(l *logger.Logger) {
	appLogger = l
}

// handleError обрабатывает ошибки и возвращает безопасное сообщение для клиента
func handleError(err error) (int, string) {
	if err == nil {
		return http.StatusOK, ""
	}

	// Ошибки валидации - возвращаем как есть
	// (пока нет специфичных доменных ошибок)

	// Ошибки БД - не раскрываем детали
	// Проверяем на типичные ошибки БД
	errStr := err.Error()
	if errors.Is(err, sql.ErrNoRows) {
		return http.StatusNotFound, "Ресурс не найден"
	}

	// Проверяем на ошибки подключения к БД
	if strings.Contains(errStr, "connection") ||
		strings.Contains(errStr, "database") ||
		strings.Contains(errStr, "sql:") ||
		strings.Contains(errStr, "pq:") {
		// Логируем реальную ошибку
		if appLogger != nil {
			appLogger.Error("Database error", map[string]interface{}{
				"error": err.Error(),
			})
		}
		return http.StatusInternalServerError, "Произошла ошибка базы данных"
	}

	// Ошибки валидации из application слоя - возвращаем как есть
	// (они уже безопасны для клиента и на русском языке)
	errMsg := err.Error()
	if len(errMsg) > 0 {
		// Ошибки валидации обычно содержат ключевые слова на русском
		if containsAny(errMsg, []string{
			"обязателен", "обязательно", "обязательна",
			"слишком длинн", "максимум",
			"должен содержать", "может содержать",
			"не может", "неверный формат",
			"формат", "символов", "цифр",
		}) {
			return http.StatusBadRequest, errMsg
		}
	}

	// Все остальные ошибки - не раскрываем детали
	if appLogger != nil {
		appLogger.Error("Internal error", map[string]interface{}{
			"error": err.Error(),
		})
	}
	return http.StatusInternalServerError, "Произошла внутренняя ошибка сервера"
}

// containsAny проверяет, содержит ли строка любую из подстрок
func containsAny(s string, substrings []string) bool {
	for _, substr := range substrings {
		if len(s) >= len(substr) {
			for i := 0; i <= len(s)-len(substr); i++ {
				if s[i:i+len(substr)] == substr {
					return true
				}
			}
		}
	}
	return false
}
