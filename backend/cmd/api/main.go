package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"ecolife-06/backend/internal/application"
	"ecolife-06/backend/internal/infrastructure/postgres"
	"ecolife-06/backend/internal/infrastructure/redis"
	"ecolife-06/backend/internal/presentation"

	"github.com/joho/godotenv"
)

func main() {
	// Загружаем переменные окружения
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Конфигурация базы данных
	dbConfig := postgres.Config{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnvInt("DB_PORT", 5432),
		User:     getEnv("DB_USER", "postgres"),
		Password: getEnv("DB_PASSWORD", "postgres"),
		DBName:   getEnv("DB_NAME", "ecolife"),
		SSLMode:  getEnv("DB_SSLMODE", "disable"),
	}

	// Подключение к PostgreSQL
	db, err := postgres.NewDB(dbConfig)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Конфигурация Redis
	redisConfig := redis.Config{
		Host:     getEnv("REDIS_HOST", "localhost"),
		Port:     getEnvInt("REDIS_PORT", 6379),
		Password: getEnv("REDIS_PASSWORD", ""),
		DB:       getEnvInt("REDIS_DB", 0),
	}

	// Подключение к Redis
	redisClient, err := redis.NewClient(redisConfig)
	if err != nil {
		log.Printf("Failed to connect to Redis: %v (continuing without Redis)", err)
	} else {
		defer redisClient.Close()
	}

	// Инициализация репозиториев
	appRepo := postgres.NewApplicationRepository(db)
	containerRepo := postgres.NewContainerRepository(db)
	benefitRepo := postgres.NewBenefitRepository(db)
	tariffRepo := postgres.NewTariffRepository(db)

	// Инициализация сервисов
	appService := application.NewApplicationService(appRepo)
	containerService := application.NewContainerService(containerRepo)
	benefitService := application.NewBenefitService(benefitRepo)
	tariffService := application.NewTariffService(tariffRepo)

	// Настройка роутера
	// Поддерживаем несколько origins через запятую
	corsOrigins := getEnv("CORS_ALLOWED_ORIGINS", "http://localhost:3000")
	allowedOrigins := parseCORSOrigins(corsOrigins)
	router := presentation.SetupRouter(
		appService,
		containerService,
		benefitService,
		tariffService,
		allowedOrigins,
	)

	// Настройка сервера
	port := getEnv("SERVER_PORT", "8080")
	host := getEnv("SERVER_HOST", "0.0.0.0")
	addr := fmt.Sprintf("%s:%s", host, port)

	srv := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	// Graceful shutdown
	go func() {
		log.Printf("Server starting on %s", addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Ожидание сигнала для graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited")
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		var result int
		if _, err := fmt.Sscanf(value, "%d", &result); err == nil {
			return result
		}
	}
	return defaultValue
}

// parseCORSOrigins парсит строку с origins, разделенными запятой
func parseCORSOrigins(origins string) []string {
	if origins == "" {
		return []string{"http://localhost:3000"}
	}

	var result []string
	parts := strings.Split(origins, ",")
	for _, part := range parts {
		trimmed := strings.TrimSpace(part)
		if trimmed != "" {
			result = append(result, trimmed)
		}
	}

	if len(result) == 0 {
		return []string{"http://localhost:3000"}
	}

	return result
}
