package usecases

import "github.com/rahulkhairwar/where-got-climbing-gym/entities"

type CoordinatesRepository interface {
	GetAllCoordinates() ([]entities.Coordinate, error)
}
