package loader

import (
	"app/internal/domain"
	"errors"
)

var (
	// ErrLoaderVehicleInternal is returned when an internal error occurs.
	ErrLoaderVehicleInternal = errors.New("loader: internal error")
)

// LoaderVehicle is the interface that wraps the basic methods for a vehicle loader.
type LoaderVehicle interface {
	Load() (v map[int]*domain.VehicleAttributes, err error)
}