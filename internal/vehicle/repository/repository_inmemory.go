package repository

import (
	"app/internal/domain"
	"fmt"
)

func NewRepositoryVehicleInMemory(db map[int]*domain.VehicleAttributes) *RepositoryVehicleInMemory {
	return &RepositoryVehicleInMemory{db: db}
}

// RepositoryVehicleInMemory is an struct that represents a vehicle storage in memory.
type RepositoryVehicleInMemory struct {
	// db is the database of vehicles.
	db map[int]*domain.VehicleAttributes
}

// GetAll returns all vehicles
func (s *RepositoryVehicleInMemory) GetAll() (v []*domain.Vehicle, err error) {
	// check if the database is empty
	if len(s.db) == 0 {
		err = ErrRepositoryVehicleNotFound
		return
	}

	// get all vehicles from the database
	v = make([]*domain.Vehicle, 0, len(s.db))
	for key, value := range s.db {
		v = append(v, &domain.Vehicle{
			Id:         key,
			Attributes: *value,
		})
	}

	return
}

// AddVehicle returns a new vehicles
func (s *RepositoryVehicleInMemory) AddVehicle(v *domain.Vehicle) (vehicle *domain.Vehicle, err error) {
	if vcl := s.db[v.Id]; vcl != nil {
		err = ErrRepositoryVehicleExist
		return
	}
	s.db[v.Id] = &v.Attributes
	fmt.Println("Se agrego a la bd correctamente")
	vehicle = v
	return
}

func (s *RepositoryVehicleInMemory) AddVehicles(vehicles []*domain.Vehicle) (v []*domain.Vehicle, err error) {
	for _, vehicle := range vehicles {
		if vcl := s.db[vehicle.Id]; vcl != nil {
			err = ErrRepositoryVehicleExist
			return
		}
	}
	for _, vehicle := range vehicles {
		addVehicle, _ := s.AddVehicle(vehicle)
		v = append(v, addVehicle)
	}
	fmt.Println("Se agregegaron los autos a la bd correctamente")
	return
}

func (s *RepositoryVehicleInMemory) GetByColorAndYear(color string, year int) (v []*domain.Vehicle, err error) {
	// get all vehicles from the database
	vehicles, err := s.GetAll()
	if err != nil {
		return
	}
	for _, vehicle := range vehicles {
		if vehicle.Attributes.Color == color && vehicle.Attributes.Year == year {
			v = append(v, vehicle)
		}
	}
	if len(v) == 0 {
		err = ErrRepositoryVehicleNotFoundWithValue
		return
	}
	return
}

func (s *RepositoryVehicleInMemory) GetByBrandAndPeriod(brand string, start int, end int) (v []*domain.Vehicle, err error) {
	// get all vehicles from the database
	vehicles, err := s.GetAll()
	if err != nil {
		return
	}
	for _, vehicle := range vehicles {
		if vehicle.Attributes.Brand == brand {
			if vehicle.Attributes.Year >= start && vehicle.Attributes.Year < end {
				v = append(v, vehicle)
			}
		}
	}
	if len(v) == 0 {
		err = ErrRepositoryVehicleNotFoundWithValue
		return
	}
	return
}

// AddVehicle returns a new vehicles
func (s *RepositoryVehicleInMemory) UpdateSpeed(v *domain.Vehicle) (vehicle *domain.Vehicle, err error) {
	if vcl := s.db[v.Id]; vcl == nil {
		err = ErrRepositoryVehicleNotFound
		return
	}
	if v.Attributes.MaxSpeed < 0 || v.Attributes.MaxSpeed > 400 {
		err = ErrRepositoryImposibleMaxSpeed
	}
	s.db[v.Id].MaxSpeed = v.Attributes.MaxSpeed
	fmt.Println("Se actualizo la velocidad correctamente")
	vehicle = &domain.Vehicle{
		Id:         v.Id,
		Attributes: *s.db[v.Id],
	}
	return
}

func (s *RepositoryVehicleInMemory) GetSpeedAverageByBrand(brand string) (average float64, err error) {
	// get all vehicles from the database
	vehicles, err := s.GetAll()
	if err != nil {
		return
	}
	var counter int
	var sumSpeed int
	for _, vehicle := range vehicles {
		if vehicle.Attributes.Brand == brand {
			counter++
			sumSpeed += vehicle.Attributes.MaxSpeed
		}
	}
	if counter == 0 {
		err = ErrRepositoryVehicleNotFoundWithValue
		return
	}
	average = float64(sumSpeed) / float64(counter)
	return
}

func (s *RepositoryVehicleInMemory) GetByFuelType(fuel string) (v []*domain.Vehicle, err error) {
	// get all vehicles from the database
	vehicles, err := s.GetAll()
	if err != nil {
		return
	}
	for _, vehicle := range vehicles {
		if vehicle.Attributes.FuelType == fuel {
			v = append(v, vehicle)
		}
	}
	if len(v) == 0 {
		err = ErrRepositoryVehicleNotFoundWithValue
		return
	}
	return
}

func (s *RepositoryVehicleInMemory) GetByWeight(min float64, max float64) (v []*domain.Vehicle, err error) {
	// get all vehicles from the database
	vehicles, err := s.GetAll()
	if err != nil {
		return
	}
	for _, vehicle := range vehicles {
		if vehicle.Attributes.Weight >= min && vehicle.Attributes.Weight <= max {
			v = append(v, vehicle)
		}
	}
	if len(v) == 0 {
		err = ErrRepositoryVehicleNotFoundWithValue
		return
	}
	return
}

func (s *RepositoryVehicleInMemory) GetById(id int) (v *domain.Vehicle, err error) {
	vehicle := s.db[id]
	if vehicle == nil {
		err = ErrRepositoryVehicleNotFound
	}
	v = &domain.Vehicle{
		Id:         id,
		Attributes: *vehicle,
	}
	return
}

func (s *RepositoryVehicleInMemory) DeleteVehicle(id int) (v *domain.Vehicle, err error) {
	// get all vehicles from the database
	v, err = s.GetById(id)
	if err != nil {
		return
	}
	delete(s.db, id)
	fmt.Println("Se elimino el vehiculo correctamente")
	return
}
