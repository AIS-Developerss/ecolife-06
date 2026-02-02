package application

import (
	"context"
	"ecolife-06/backend/internal/domain"
	"errors"
)

// ApplicationService представляет сервис для работы с заявками
type ApplicationService struct {
	appRepo domain.ApplicationRepository
}

// NewApplicationService создает новый сервис заявок
func NewApplicationService(appRepo domain.ApplicationRepository) *ApplicationService {
	return &ApplicationService{
		appRepo: appRepo,
	}
}

// CreateApplication создает новую заявку
func (s *ApplicationService) CreateApplication(ctx context.Context, req *CreateApplicationRequest) (*domain.Application, error) {
	if req.FullName == "" {
		return nil, errors.New("full name is required")
	}
	if req.Phone == "" {
		return nil, errors.New("phone is required")
	}
	if req.Address == "" {
		return nil, errors.New("address is required")
	}
	if req.District == "" {
		return nil, errors.New("district is required")
	}

	app := &domain.Application{
		ID:          generateID(),
		FullName:    req.FullName,
		Phone:       req.Phone,
		Address:     req.Address,
		District:    req.District,
		ContainerID: req.ContainerID,
		ServiceType: req.ServiceType,
		Status:      string(domain.StatusNew),
	}

	if err := s.appRepo.Create(ctx, app); err != nil {
		return nil, err
	}

	return app, nil
}

// GetApplication получает заявку по ID
func (s *ApplicationService) GetApplication(ctx context.Context, id string) (*domain.Application, error) {
	return s.appRepo.GetByID(ctx, id)
}

// GetAllApplications получает все заявки с пагинацией
func (s *ApplicationService) GetAllApplications(ctx context.Context, limit, offset int) ([]*domain.Application, error) {
	return s.appRepo.GetAll(ctx, limit, offset)
}

// UpdateApplicationStatus обновляет статус заявки
func (s *ApplicationService) UpdateApplicationStatus(ctx context.Context, id string, status domain.ApplicationStatus) error {
	app, err := s.appRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	app.Status = string(status)
	return s.appRepo.Update(ctx, app)
}

// CreateApplicationRequest представляет запрос на создание заявки
type CreateApplicationRequest struct {
	FullName    string `json:"full_name"`
	Phone       string `json:"phone"`
	Address     string `json:"address"`
	District    string `json:"district"`
	ContainerID string `json:"container_id,omitempty"`
	ServiceType string `json:"service_type"`
}
