package domain

import (
	"context"
)

// ApplicationRepository определяет интерфейс для работы с заявками
type ApplicationRepository interface {
	Create(ctx context.Context, application *Application) error
}

