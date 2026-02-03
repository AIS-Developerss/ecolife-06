package domain

import (
	"time"
)

// Application представляет заявку на услугу
type Application struct {
	ID          string    `json:"id"`
	FullName    string    `json:"full_name"`
	Phone       string    `json:"phone"`
	Address     string    `json:"address"`
	District    string    `json:"district"`
	ContainerID string    `json:"container_id,omitempty"`
	ServiceType string    `json:"service_type"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

