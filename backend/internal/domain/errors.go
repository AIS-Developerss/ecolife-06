package domain

import "errors"

var (
	ErrApplicationNotFound = errors.New("application not found")
	ErrContainerNotFound   = errors.New("container not found")
	ErrBenefitNotFound     = errors.New("benefit not found")
	ErrTariffNotFound      = errors.New("tariff not found")
)

