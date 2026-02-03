package postgres

import (
	"context"
	"database/sql"
	"ecolife-06/backend/internal/domain"
	"time"
)

// ApplicationRepository реализует domain.ApplicationRepository
type ApplicationRepository struct {
	db *sql.DB
}

// NewApplicationRepository создает новый репозиторий заявок
func NewApplicationRepository(db *sql.DB) *ApplicationRepository {
	return &ApplicationRepository{db: db}
}

// Create создает новую заявку
func (r *ApplicationRepository) Create(ctx context.Context, app *domain.Application) error {
	query := `
		INSERT INTO applications (id, full_name, phone, address, district, container_id, service_type, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`
	now := time.Now()
	app.CreatedAt = now
	app.UpdatedAt = now

	_, err := r.db.ExecContext(ctx, query,
		app.ID, app.FullName, app.Phone, app.Address, app.District,
		app.ContainerID, app.ServiceType, app.Status, app.CreatedAt, app.UpdatedAt,
	)
	return err
}

