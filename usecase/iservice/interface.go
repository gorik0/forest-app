package iservice

import "awesomeProject/domain"

type ForestService interface {
	GetAnimals() ([]*domain.Animal, error)
	GetSquare() (float64, error)
	PopulateWithAnimals(animal *domain.Animal) (bool, error)
}
