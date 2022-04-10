package usecases

import "github.com/rahulkhairwar/where-got-climbing-gym/entities"

func GetCoordinates(repo CoordinatesRepository) ([]entities.Coordinate, error) {
	coordinates, err := repo.GetAllCoordinates()
	if err != nil {
		return nil, ErrInternal
	}

	return coordinates, nil
}
