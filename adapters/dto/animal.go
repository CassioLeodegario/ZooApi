package dto

import "github.com/cassioleodegario/zooapi/domain"

type Animal struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Species *Species `json:"specie"`
	Diet    string   `json:"diet"`
}

// Species Model
type Species struct {
	Name          string `json:"name"`
	Family        string `json:"family"`
	Domestication string `json:"domestication"`
	Endangered    string `json:"endangered"`
}

func NewAnimal() *Animal {
	return &Animal{}
}

func (a *Animal) Bind(animal *domain.Animal) (*domain.Animal, error) {
	if a.ID != "" {
		animal.ID = a.ID
	}
	animal.Name = a.Name
	animal.Diet = a.Diet
	animal.Species = (*domain.Species)(a.Species)
	_, err := animal.IsValid()
	if err != nil {
		return &domain.Animal{}, err
	}
	return animal, nil
}
