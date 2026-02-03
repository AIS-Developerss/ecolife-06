package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// RateLimiter представляет простой rate limiter на основе in-memory хранилища
type RateLimiter struct {
	requests map[string][]time.Time
	mu       sync.RWMutex
	limit    int           // Максимальное количество запросов
	window   time.Duration // Временное окно
}

// NewRateLimiter создает новый rate limiter
func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
	rl := &RateLimiter{
		requests: make(map[string][]time.Time),
		limit:    limit,
		window:   window,
	}

	// Очистка старых записей каждую минуту
	go rl.cleanup()

	return rl
}

// cleanup удаляет старые записи из map
func (rl *RateLimiter) cleanup() {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		rl.mu.Lock()
		now := time.Now()
		for ip, times := range rl.requests {
			// Удаляем записи старше window
			validTimes := []time.Time{}
			for _, t := range times {
				if now.Sub(t) < rl.window {
					validTimes = append(validTimes, t)
				}
			}
			if len(validTimes) == 0 {
				delete(rl.requests, ip)
			} else {
				rl.requests[ip] = validTimes
			}
		}
		rl.mu.Unlock()
	}
}

// Allow проверяет, разрешен ли запрос
func (rl *RateLimiter) Allow(ip string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	times := rl.requests[ip]

	// Удаляем старые записи
	validTimes := []time.Time{}
	for _, t := range times {
		if now.Sub(t) < rl.window {
			validTimes = append(validTimes, t)
		}
	}

	// Проверяем лимит
	if len(validTimes) >= rl.limit {
		return false
	}

	// Добавляем новую запись
	validTimes = append(validTimes, now)
	rl.requests[ip] = validTimes

	return true
}

// RateLimitMiddleware создает middleware для rate limiting
func RateLimitMiddleware(limit int, window time.Duration) gin.HandlerFunc {
	limiter := NewRateLimiter(limit, window)

	return func(c *gin.Context) {
		// Получаем IP адрес клиента
		ip := c.ClientIP()

		if !limiter.Allow(ip) {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"message": "Too many requests. Please try again later.",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

