package postgres

import (
	"context"
	"database/sql"
	"ecolife-06/backend/internal/domain"
)

// BenefitRepository реализует domain.BenefitRepository
type BenefitRepository struct {
	db *sql.DB
}

// NewBenefitRepository создает новый репозиторий льгот
func NewBenefitRepository(db *sql.DB) *BenefitRepository {
	return &BenefitRepository{db: db}
}

// GetAll получает все льготы
func (r *BenefitRepository) GetAll(ctx context.Context) ([]*domain.Benefit, error) {
	query := `
		SELECT id, category, description, discount, is_active
		FROM benefits
		WHERE is_active = true
		ORDER BY discount DESC
	`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var benefits []*domain.Benefit
	for rows.Next() {
		benefit := &domain.Benefit{}
		err := rows.Scan(
			&benefit.ID, &benefit.Category, &benefit.Description, &benefit.Discount, &benefit.IsActive,
		)
		if err != nil {
			return nil, err
		}
		benefits = append(benefits, benefit)
	}
	return benefits, rows.Err()
}

// GetByCategory получает льготу по категории
func (r *BenefitRepository) GetByCategory(ctx context.Context, category string) (*domain.Benefit, error) {
	query := `
		SELECT id, category, description, discount, is_active
		FROM benefits
		WHERE category = $1 AND is_active = true
	`
	benefit := &domain.Benefit{}
	err := r.db.QueryRowContext(ctx, query, category).Scan(
		&benefit.ID, &benefit.Category, &benefit.Description, &benefit.Discount, &benefit.IsActive,
	)
	if err == sql.ErrNoRows {
		return nil, domain.ErrBenefitNotFound
	}
	if err != nil {
		return nil, err
	}
	return benefit, nil
}

