package service

import (
	"app/internal/domain"
	"app/internal/vehicle/repository"
	"errors"
	"fmt"
)

// ServiceVehicleDefault is an struct that represents a vehicle service.
type ServiceVehicleDefault struct {
	rp repository.RepositoryVehicle
}

// NewServiceVehicleDefault returns a new instance of a vehicle service.
func NewServiceVehicleDefault(rp repository.RepositoryVehicle) *ServiceVehicleDefault {
	return &ServiceVehicleDefault{rp: rp}
}

func validateErrors(error error) (err error) {
	switch {
	case errors.Is(error, repository.ErrRepositoryVehicleNotFound):
		return fmt.Errorf("%w. %v", ErrServiceVehicleNotFound, err)
	case errors.Is(error, repository.ErrRepositoryVehicleExist):
		return fmt.Errorf("%w. %v", ErrServiceVehicleExist, err)
	case errors.Is(error, repository.ErrRepositoryVehicleNotFoundWithValue):
		return fmt.Errorf("%w. %v", ErrServiceVehicleNotFoundWithValue, err)
	case errors.Is(error, repository.ErrRepositoryImposibleMaxSpeed):
		return fmt.Errorf("%w. %v", ErrServiceImposibleMaxSpeed, err)
	default:
		return fmt.Errorf("%w. %v", ErrServiceVehicleInternal, err)
	}
}

// GetAll returns all vehicles.
func (s *ServiceVehicleDefault) GetAll() (v []*domain.Vehicle, err error) {
	v, err = s.rp.GetAll()
	if err != nil {
		err = validateErrors(err)
		return
	}
	return
}

// AddVehicle add a new vehicle.
func (s *ServiceVehicleDefault) AddVehicle(vehicle *domain.Vehicle) (v *domain.Vehicle, err error) {
	v, err = s.rp.AddVehicle(vehicle)
	if err != nil {
		err = validateErrors(err)
		return
	}

	return
}

func (s *ServiceVehicleDefault) GetByColorAndYear(color string, year int) (v []*domain.Vehicle, err error) {
	v, err = s.rp.GetByColorAndYear(color, year)
	if err != nil {
		err = validateErrors(err)
		return
	}

	return
}

func (s *ServiceVehicleDefault) GetByBrandAndPeriod(brand string, start int, end int) (v []*domain.Vehicle, err error) {
	v, err = s.rp.GetByBrandAndPeriod(brand, start, end)
	if err != nil {
		err = validateErrors(err)
		return
	}
	return
}

func (s *ServiceVehicleDefault) UpdateSpeed(vehicle *domain.Vehicle) (v *domain.Vehicle, err error) {
	v, err = s.rp.UpdateSpeed(vehicle)
	if err != nil {
		err = validateErrors(err)
		return
	}

	return
}

func (s *ServiceVehicleDefault) GetSpeedAverageByBrand(brand string) (average float64, err error) {
	average, err = s.rp.GetSpeedAverageByBrand(brand)
	if err != nil {
		err = validateErrors(err)
		return
	}
	return
}

func (s *ServiceVehicleDefault) GetByFuelType(fuel string) (v []*domain.Vehicle, err error) {
	v, err = s.rp.GetByFuelType(fuel)
	if err != nil {
		err = validateErrors(err)
		return
	}
	return
}

func (s *ServiceVehicleDefault) GetByWeight(min float64, max float64) (v []*domain.Vehicle, err error) {
	v, err = s.rp.GetByWeight(min, max)
	if err != nil {
		err = validateErrors(err)
		return
	}
	return
}

func (s *ServiceVehicleDefault) DeleteVehicle(id int) (v *domain.Vehicle, err error) {
	v, err = s.rp.DeleteVehicle(id)
	if err != nil {
		err = validateErrors(err)
		return
	}
	return
}

func (s *ServiceVehicleDefault) AddVehicles(vehicles []*domain.Vehicle) (v []*domain.Vehicle, err error) {
	v, err = s.rp.AddVehicles(vehicles)
	if err != nil {
		err = validateErrors(err)
		return
	}

	return
}
