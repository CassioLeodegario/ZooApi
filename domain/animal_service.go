package domain

type AnimalService struct {
	Persistence AnimalPersistenceInterface
}

func NewAnimalService(persistence AnimalPersistenceInterface) *AnimalService {
	return &AnimalService{Persistence: persistence}
}

func (a *AnimalService) Get(id string) (AnimalInterface, error) {
	animal, err := a.Persistence.Get(id)
	if err != nil {
		return nil, err
	}

	return animal, nil
}

func (a *AnimalService) Create(animal AnimalInterface) (AnimalInterface, error) {
	animal, err := a.Persistence.Save(animal)
	if err != nil {
		return nil, err
	}

	return animal, nil
}

func (a *AnimalService) Update(animal AnimalInterface) (AnimalInterface, error) {
	animal, err := a.Persistence.Save(animal)
	if err != nil {
		return nil, err
	}

	return animal, nil
}

func (a *AnimalService) Delete(animal AnimalInterface) (AnimalInterface, error) {
	animal, err := a.Persistence.Save(animal)
	if err != nil {
		return nil, err
	}

	return animal, nil
}
