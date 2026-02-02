package postgres

import (
	"context"
	"database/sql"
	"ecolife-06/backend/internal/domain"
	"time"
)

// TariffRepository реализует domain.TariffRepository
type TariffRepository struct {
	db *sql.DB
}

// NewTariffRepository создает новый репозиторий тарифов
func NewTariffRepository(db *sql.DB) *TariffRepository {
	return &TariffRepository{db: db}
}

// GetCurrent получает текущий действующий тариф
func (r *TariffRepository) GetCurrent(ctx context.Context) (*domain.Tariff, error) {
	query := `
		SELECT id, price, valid_from, valid_to, is_active, description
		FROM tariffs
		WHERE is_active = true 
		  AND valid_from <= $1 
		  AND (valid_to IS NULL OR valid_to >= $1)
		ORDER BY valid_from DESC
		LIMIT 1
	`
	tariff := &domain.Tariff{}
	var validTo sql.NullTime
	err := r.db.QueryRowContext(ctx, query, time.Now()).Scan(
		&tariff.ID, &tariff.Price, &tariff.ValidFrom, &validTo, &tariff.IsActive, &tariff.Description,
	)
	if err == sql.ErrNoRows {
		return nil, domain.ErrTariffNotFound
	}
	if err != nil {
		return nil, err
	}
	if validTo.Valid {
		tariff.ValidTo = &validTo.Time
	}
	return tariff, nil
}

// GetAll получает все тарифы
func (r *TariffRepository) GetAll(ctx context.Context) ([]*domain.Tariff, error) {
	query := `
		SELECT id, price, valid_from, valid_to, is_active, description
		FROM tariffs
		ORDER BY valid_from DESC
	`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tariffs []*domain.Tariff
	for rows.Next() {
		tariff := &domain.Tariff{}
		var validTo sql.NullTime
		err := rows.Scan(
			&tariff.ID, &tariff.Price, &tariff.ValidFrom, &validTo, &tariff.IsActive, &tariff.Description,
		)
		if err != nil {
			return nil, err
		}
		if validTo.Valid {
			tariff.ValidTo = &validTo.Time
		}
		tariffs = append(tariffs, tariff)
	}
	return tariffs, rows.Err()
}

