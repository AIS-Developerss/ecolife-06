package application

import "github.com/google/uuid"

// generateID генерирует новый UUID
func generateID() string {
	return uuid.New().String()
}

