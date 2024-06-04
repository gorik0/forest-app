package usecase

import (
	"awesomeProject/domain"
	"awesomeProject/repository/interface"
	"awesomeProject/usecase/iservice"
)

type ForestServiceImpl struct {
	f _interface.ForestRepository
}

func (f ForestServiceImpl) GetAnimals() ([]*domain.Animal, error) {
	animals, err := f.f.GetAnimals()
	if err != nil {
		return nil, err
	}
	return animals, nil
}

func (f ForestServiceImpl) GetSquare() (float64, error) {
	square, err := f.f.GetSquare()
	if err != nil {

		return -1, err

	}
	return square, nil
}

func (f ForestServiceImpl) PopulateWithAnimals(animal *domain.Animal) (bool, error) {
	_, err := f.f.PopulateWithAnimals(animal)
	if err != nil {
		return false, err
	}
	return true, nil
}

func NewForestService(repo _interface.ForestRepository) iservice.ForestService {
	return &ForestServiceImpl{f: repo}
}
