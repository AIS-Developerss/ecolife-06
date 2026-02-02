package postgres

import (
	"context"
	"database/sql"
	"ecolife-06/backend/internal/domain"
)

// ContainerRepository реализует domain.ContainerRepository
type ContainerRepository struct {
	db *sql.DB
}

// NewContainerRepository создает новый репозиторий контейнеров
func NewContainerRepository(db *sql.DB) *ContainerRepository {
	return &ContainerRepository{db: db}
}

// GetByID получает контейнер по ID
func (r *ContainerRepository) GetByID(ctx context.Context, id string) (*domain.Container, error) {
	query := `
		SELECT id, volume, price, description, is_active
		FROM containers
		WHERE id = $1
	`
	container := &domain.Container{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&container.ID, &container.Volume, &container.Price, &container.Description, &container.IsActive,
	)
	if err == sql.ErrNoRows {
		return nil, domain.ErrContainerNotFound
	}
	if err != nil {
		return nil, err
	}
	return container, nil
}

// GetAll получает все контейнеры
func (r *ContainerRepository) GetAll(ctx context.Context) ([]*domain.Container, error) {
	query := `
		SELECT id, volume, price, description, is_active
		FROM containers
		WHERE is_active = true
		ORDER BY volume
	`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var containers []*domain.Container
	for rows.Next() {
		container := &domain.Container{}
		err := rows.Scan(
			&container.ID, &container.Volume, &container.Price, &container.Description, &container.IsActive,
		)
		if err != nil {
			return nil, err
		}
		containers = append(containers, container)
	}
	return containers, rows.Err()
}

// GetByVolume получает контейнер по объему
func (r *ContainerRepository) GetByVolume(ctx context.Context, volume int) (*domain.Container, error) {
	query := `
		SELECT id, volume, price, description, is_active
		FROM containers
		WHERE volume = $1 AND is_active = true
	`
	container := &domain.Container{}
	err := r.db.QueryRowContext(ctx, query, volume).Scan(
		&container.ID, &container.Volume, &container.Price, &container.Description, &container.IsActive,
	)
	if err == sql.ErrNoRows {
		return nil, domain.ErrContainerNotFound
	}
	if err != nil {
		return nil, err
	}
	return container, nil
}

