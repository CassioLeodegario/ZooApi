package domain

import (
	"github.com/asaskevich/govalidator"
)

func init() {

}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type AnimalInterface interface {
	IsValid() (bool, error)
	GetID() string
	GetName() string
	GetSpecie() Species
	GetDiet() string
}

type AnimalServiceInterface interface {
	GetAll() ([]AnimalInterface, error)
	Get(id string) (AnimalInterface, error)
	Create(animal AnimalInterface) (AnimalInterface, error)
	Update(animal AnimalInterface) (AnimalInterface, error)
	Delete(animal Animal) error
}

type AnimalReader interface {
	Get(id string) (AnimalInterface, error)
}

type AnimalWriter interface {
	Save(animal AnimalInterface) (AnimalInterface, error)
}

type AnimalPersistenceInterface interface {
	AnimalReader
	AnimalWriter
}

type Animal struct {
	ID      string   `valid:"uuidv4"`
	Name    string   `valid:"required"`
	Species *Species `valid:"required"`
	Diet    string   `valid:"string,optional"`
}

// Species Model
type Species struct {
	Name          string `valid:"required"`
	Family        string `valid:"required"`
	Domestication string `valid:"required"`
	Endangered    string `valid:"required"`
}

func NewAnimal() *Animal {
	return &Animal{}
}

func (a *Animal) IsValid() (bool, error) {
	_, err := govalidator.ValidateStruct(a)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (a *Animal) GetID() string {
	return a.ID
}

func (a *Animal) GetName() string {
	return a.Name
}

func (a *Animal) GetSpecie() Species {
	return *a.Species
}

func (a *Animal) GetDiet() string {
	return a.Diet
}
