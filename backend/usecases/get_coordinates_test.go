package usecases_test

import (
	"github.com/rahulkhairwar/where-got-climbing-gym/entities"
	"github.com/rahulkhairwar/where-got-climbing-gym/usecases"
	"github.com/rahulkhairwar/where-got-climbing-gym/usecases/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

var dummyCoordinates = []entities.Coordinate{
	{
		Title:       "Coordinate 1",
		Description: "Description for coordinate 1",
		X:           10,
		Y:           25,
	},
	{
		Title:       "Coordinate 2",
		Description: "Description for coordinate 2",
		X:           101.6,
		Y:           568.2,
	},
	{
		Title:       "Coordinate 3",
		Description: "Description for coordinate 3",
		X:           9.22210,
		Y:           2.288925,
	},
}

func TestGetCoordinates(t *testing.T) {
	t.Run("Returns ErrInternal when CoordinatesRepository returns error", func(t *testing.T) {
		repo := new(mocks.MockCoordinatesRepo)

		repo.On(repo.GetAllCoordinatesStr()).Return(nil, usecases.ErrInternal)

		coordinates, err := usecases.GetCoordinates(repo)
		assert.EqualError(t, err, usecases.ErrInternal.Error())
		assert.Empty(t, coordinates)
	})

	t.Run("Returns Coordinates from CoordinatesRepository", func(t *testing.T) {
		repo := new(mocks.MockCoordinatesRepo)

		repo.On(repo.GetAllCoordinatesStr()).Return(dummyCoordinates, nil)

		coordinates, err := usecases.GetCoordinates(repo)
		assert.NoError(t, err)
		assert.Equal(t, dummyCoordinates, coordinates)
	})
}
