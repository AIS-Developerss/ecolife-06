package domain

import "time"

// Tariff представляет тариф на услугу
type Tariff struct {
	ID          string     `json:"id"`
	Price       float64    `json:"price"` // цена за куб.м
	ValidFrom   time.Time  `json:"valid_from"`
	ValidTo     *time.Time `json:"valid_to,omitempty"`
	IsActive    bool       `json:"is_active"`
	Description string     `json:"description"`
}

