package service

import (
	"app/internal/domain"
	"errors"
)

// ServiceVehicle is the interface that wraps the basic methods for a vehicle service.
// - conections with external apis
// - business logic
type ServiceVehicle interface {
	// GetAll returns all vehicles
	GetAll() (v []*domain.Vehicle, err error)
	GetByColorAndYear(color string, year int) (v []*domain.Vehicle, err error)
	GetByBrandAndPeriod(brand string, start int, end int) (v []*domain.Vehicle, err error)
	GetSpeedAverageByBrand(brand string) (v float64, err error)
	GetByFuelType(fuel string) (v []*domain.Vehicle, err error)
	GetByWeight(min float64, max float64) (v []*domain.Vehicle, err error)

	AddVehicle(attributes *domain.Vehicle) (v *domain.Vehicle, err error)
	AddVehicles(vehicles []*domain.Vehicle) (v []*domain.Vehicle, err error)

	UpdateSpeed(vehicle *domain.Vehicle) (v *domain.Vehicle, err error)

	DeleteVehicle(id int) (v *domain.Vehicle, err error)
}

var (
	// ErrServiceVehicleInternal is returned when an internal error occurs.
	ErrServiceVehicleInternal = errors.New("service: internal error")

	// ErrServiceVehicleNotFound is returned when no vehicle is found.
	ErrServiceVehicleNotFound = errors.New("service: vehicle not found")

	ErrServiceVehicleExist             = errors.New("service: identificador del vehículo ya existente")
	ErrServiceVehicleNotFoundWithValue = errors.New("service: no se encontraron vehiculos con esos criterios")
	ErrServiceImposibleMaxSpeed        = errors.New("service: Velocidad mal formada o fuera de rango")
)
