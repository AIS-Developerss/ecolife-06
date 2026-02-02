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

// GetByID получает заявку по ID
func (r *ApplicationRepository) GetByID(ctx context.Context, id string) (*domain.Application, error) {
	query := `
		SELECT id, full_name, phone, address, district, container_id, service_type, status, created_at, updated_at
		FROM applications
		WHERE id = $1
	`
	app := &domain.Application{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&app.ID, &app.FullName, &app.Phone, &app.Address, &app.District,
		&app.ContainerID, &app.ServiceType, &app.Status, &app.CreatedAt, &app.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, domain.ErrApplicationNotFound
	}
	if err != nil {
		return nil, err
	}
	return app, nil
}

// GetAll получает все заявки с пагинацией
func (r *ApplicationRepository) GetAll(ctx context.Context, limit, offset int) ([]*domain.Application, error) {
	query := `
		SELECT id, full_name, phone, address, district, container_id, service_type, status, created_at, updated_at
		FROM applications
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`
	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var applications []*domain.Application
	for rows.Next() {
		app := &domain.Application{}
		err := rows.Scan(
			&app.ID, &app.FullName, &app.Phone, &app.Address, &app.District,
			&app.ContainerID, &app.ServiceType, &app.Status, &app.CreatedAt, &app.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		applications = append(applications, app)
	}
	return applications, rows.Err()
}

// Update обновляет заявку
func (r *ApplicationRepository) Update(ctx context.Context, app *domain.Application) error {
	query := `
		UPDATE applications
		SET full_name = $2, phone = $3, address = $4, district = $5, container_id = $6, 
		    service_type = $7, status = $8, updated_at = $9
		WHERE id = $1
	`
	app.UpdatedAt = time.Now()
	_, err := r.db.ExecContext(ctx, query,
		app.ID, app.FullName, app.Phone, app.Address, app.District,
		app.ContainerID, app.ServiceType, app.Status, app.UpdatedAt,
	)
	return err
}

// Delete удаляет заявку
func (r *ApplicationRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM applications WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

