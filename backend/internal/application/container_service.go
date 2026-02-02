package application

import (
	"context"
	"ecolife-06/backend/internal/domain"
)

// ContainerService представляет сервис для работы с контейнерами
type ContainerService struct {
	containerRepo domain.ContainerRepository
}

// NewContainerService создает новый сервис контейнеров
func NewContainerService(containerRepo domain.ContainerRepository) *ContainerService {
	return &ContainerService{
		containerRepo: containerRepo,
	}
}

// GetAllContainers получает все доступные контейнеры
func (s *ContainerService) GetAllContainers(ctx context.Context) ([]*domain.Container, error) {
	return s.containerRepo.GetAll(ctx)
}

// GetContainerByID получает контейнер по ID
func (s *ContainerService) GetContainerByID(ctx context.Context, id string) (*domain.Container, error) {
	return s.containerRepo.GetByID(ctx, id)
}

// GetContainerByVolume получает контейнер по объему
func (s *ContainerService) GetContainerByVolume(ctx context.Context, volume int) (*domain.Container, error) {
	return s.containerRepo.GetByVolume(ctx, volume)
}

