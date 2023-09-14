package repository

import (
	"app/internal/domain"
	"errors"
)

// RepositoryVehicle is the interface that wraps the basic methods for a vehicle repository.
type RepositoryVehicle interface {
	// GetAll returns all vehicles
	GetAll() (v []*domain.Vehicle, err error)
	GetByColorAndYear(color string, year int) (v []*domain.Vehicle, err error)
	GetByBrandAndPeriod(brand string, start int, end int) (v []*domain.Vehicle, err error)
	GetSpeedAverageByBrand(brand string) (v float64, err error)
	GetByFuelType(fuel string) (v []*domain.Vehicle, err error)
	GetByWeight(min float64, max float64) (v []*domain.Vehicle, err error)

	AddVehicle(attributes *domain.Vehicle) (v *domain.Vehicle, err error)
	AddVehicles(attributes []*domain.Vehicle) (v []*domain.Vehicle, err error)

	UpdateSpeed(vehicle *domain.Vehicle) (v *domain.Vehicle, err error)

	DeleteVehicle(id int) (v *domain.Vehicle, err error)
}

var (
	// ErrRepositoryVehicleInternal is returned when an internal error occurs.
	ErrRepositoryVehicleInternal = errors.New("repository: internal error")

	// ErrRepositoryVehicleNotFound is returned when a vehicle is not found.
	ErrRepositoryVehicleNotFound = errors.New("repository: vehicle not found")

	ErrRepositoryVehicleExist             = errors.New("repository: identificador del vehículo ya existente")
	ErrRepositoryVehicleNotFoundWithValue = errors.New("repository: vehiculos no encontrados con esos criterios")
	ErrRepositoryImposibleMaxSpeed        = errors.New("repository: Velocidad mal formada o fuera de rango")
)
