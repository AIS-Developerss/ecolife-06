package application

import (
	"context"
	"ecolife-06/backend/internal/domain"
)

// TariffService представляет сервис для работы с тарифами
type TariffService struct {
	tariffRepo domain.TariffRepository
}

// NewTariffService создает новый сервис тарифов
func NewTariffService(tariffRepo domain.TariffRepository) *TariffService {
	return &TariffService{
		tariffRepo: tariffRepo,
	}
}

// GetCurrentTariff получает текущий действующий тариф
func (s *TariffService) GetCurrentTariff(ctx context.Context) (*domain.Tariff, error) {
	return s.tariffRepo.GetCurrent(ctx)
}

// GetAllTariffs получает все тарифы
func (s *TariffService) GetAllTariffs(ctx context.Context) ([]*domain.Tariff, error) {
	return s.tariffRepo.GetAll(ctx)
}

