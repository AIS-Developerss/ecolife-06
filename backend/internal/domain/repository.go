package domain

import (
	"context"
)

// ApplicationRepository определяет интерфейс для работы с заявками
type ApplicationRepository interface {
	Create(ctx context.Context, application *Application) error
	GetByID(ctx context.Context, id string) (*Application, error)
	GetAll(ctx context.Context, limit, offset int) ([]*Application, error)
	Update(ctx context.Context, application *Application) error
	Delete(ctx context.Context, id string) error
}

// ContainerRepository определяет интерфейс для работы с контейнерами
type ContainerRepository interface {
	GetByID(ctx context.Context, id string) (*Container, error)
	GetAll(ctx context.Context) ([]*Container, error)
	GetByVolume(ctx context.Context, volume int) (*Container, error)
}

// BenefitRepository определяет интерфейс для работы с льготами
type BenefitRepository interface {
	GetAll(ctx context.Context) ([]*Benefit, error)
	GetByCategory(ctx context.Context, category string) (*Benefit, error)
}

// TariffRepository определяет интерфейс для работы с тарифами
type TariffRepository interface {
	GetCurrent(ctx context.Context) (*Tariff, error)
	GetAll(ctx context.Context) ([]*Tariff, error)
}

