package application

import (
	"context"
	"ecolife-06/backend/internal/domain"
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

// CreateFeedbackRequest представляет запрос на обратную связь (минимальные поля)
type CreateFeedbackRequest struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

// CreateFeedback создает заявку из формы обратной связи (только имя и телефон)
func (s *ApplicationService) CreateFeedback(ctx context.Context, req *CreateFeedbackRequest) (*domain.Application, error) {
	// Валидация имени
	if err := ValidateName(req.Name); err != nil {
		return nil, err
	}
	
	// Валидация телефона
	if err := ValidatePhone(req.Phone); err != nil {
		return nil, err
	}
	
	// Санитизация данных
	name := SanitizeString(req.Name)
	phone := SanitizeString(req.Phone)

	app := &domain.Application{
		ID:          generateID(),
		FullName:    name,
		Phone:       phone,
		Address:     "", // Необязательное поле для формы обратной связи
		District:    "", // Необязательное поле для формы обратной связи
		ContainerID: "",
		ServiceType: "household", // Значение по умолчанию
		Status:      "new",
	}

	if err := s.appRepo.Create(ctx, app); err != nil {
		return nil, err
	}

	return app, nil
}
