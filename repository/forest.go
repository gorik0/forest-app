package repository

import (
	"awesomeProject/domain"
	_interface "awesomeProject/repository/interface"
	"database/sql"
)

type ForestRepositoryImpl struct {
	DB *sql.DB
}

func (f ForestRepositoryImpl) GetAnimals() ([]*domain.Animal, error) {
	return nil, nil
}

func (f ForestRepositoryImpl) GetSquare() (float64, error) {

	return 123, nil

}

func (f ForestRepositoryImpl) PopulateWithAnimals(animal *domain.Animal) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func NewForestRepo(db *sql.DB) _interface.ForestRepository {
	return &ForestRepositoryImpl{
		DB: db,
	}
}
