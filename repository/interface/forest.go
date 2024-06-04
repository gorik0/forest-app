package _interface

import "awesomeProject/domain"

type ForestRepository interface {
	GetAnimals() ([]*domain.Animal, error)
	GetSquare() (float64, error)
	PopulateWithAnimals(animal *domain.Animal) (bool, error)
}
