package vehicle

import "github.com/google/uuid"

type Repository interface {
	Save(vehicle Vehicle) error
	FindByID(id uuid.UUID) (Vehicle, error)
	FindAll() ([]Vehicle, error)
}

type InMemoryRepository struct {
	vehicles []Vehicle
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		vehicles: []Vehicle{},
	}
}

// FindAll implements Repository.
func (i *InMemoryRepository) FindAll() ([]Vehicle, error) {
	return i.vehicles, nil
}

// FindByID implements Repository.
func (i *InMemoryRepository) FindByID(id uuid.UUID) (Vehicle, error) {
	for _, v := range i.vehicles {
		if v.ID == id {
			return v, nil
		}
	}

	return Vehicle{}, nil
}

// Save implements Repository.
func (i *InMemoryRepository) Save(vehicle Vehicle) error {
	i.vehicles = append(i.vehicles, vehicle)
	return nil
}

var _ Repository = &InMemoryRepository{}
