package application

import (
	"context"
	"ecolife-06/backend/internal/domain"
)

// BenefitService представляет сервис для работы с льготами
type BenefitService struct {
	benefitRepo domain.BenefitRepository
}

// NewBenefitService создает новый сервис льгот
func NewBenefitService(benefitRepo domain.BenefitRepository) *BenefitService {
	return &BenefitService{
		benefitRepo: benefitRepo,
	}
}

// GetAllBenefits получает все доступные льготы
func (s *BenefitService) GetAllBenefits(ctx context.Context) ([]*domain.Benefit, error) {
	return s.benefitRepo.GetAll(ctx)
}

// GetBenefitByCategory получает льготу по категории
func (s *BenefitService) GetBenefitByCategory(ctx context.Context, category string) (*domain.Benefit, error) {
	return s.benefitRepo.GetByCategory(ctx, category)
}

