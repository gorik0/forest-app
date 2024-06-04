package repository

import (
	"awesomeProject/domain"
	_interface "awesomeProject/repository/interface"
	"database/sql"
	"fmt"
	"github.com/lib/pq"
)

type ForestRepositoryImpl struct {
	DB *sql.DB
}

func (f ForestRepositoryImpl) GetAnimals() ([]*domain.Animal, error) {
	q := `select animals from  forests where id =1`
	var animalsStr []string
	var animals []*domain.Animal

	err := f.DB.QueryRow(q).Scan((*pq.StringArray)(&animalsStr))
	if err != nil {
		return nil, fmt.Errorf("error while getting(or scanning) animals sin your forest...sorry ... %d", err)
	}
	for _, animal := range animalsStr {
		animals = append(animals, &domain.Animal{Name: animal})
	}
	return animals, nil
}

func (f ForestRepositoryImpl) GetSquare() (float64, error) {
	q := `select "square" from forests where "id"=1`
	var square float64
	err := f.DB.QueryRow(q).Scan(&square)
	if err != nil {
		return -1, err
	}
	return square, nil

}

func (f ForestRepositoryImpl) PopulateWithAnimals(animal *domain.Animal) (bool, error) {
	q := `UPDATE  forests set animals=array_append(animals, $1)`

	_, err := f.DB.Exec(q, animal.Name)
	if err != nil {
		return false, err
	}
	return true, nil
}

func NewForestRepo(db *sql.DB) _interface.ForestRepository {
	return &ForestRepositoryImpl{
		DB: db,
	}
}
