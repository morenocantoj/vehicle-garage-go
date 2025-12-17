package vehicle_test

import (
	"fmt"
	"go-highschool-api/src/vehicle"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestNewCar_WhenProvidingValidData_ReturnsExpectedResult(t *testing.T) {
	// Given
	expectedBrand := "Toyota"
	expectedModel := "Corolla"
	expectedYear := 2020

	// When
	newCar := vehicle.NewCar(expectedBrand, expectedModel, expectedYear)

	// Then
	require.NotEqual(t, newCar.ID, uuid.Nil)
	require.Equal(t, expectedBrand, newCar.Brand)
	require.Equal(t, expectedModel, newCar.Model)
	require.Equal(t, expectedYear, newCar.Year)
}

func TestContaminationBadge_WhenProvidingDifferentYears_ReturnsExpectedBadge(t *testing.T) {
	tests := []struct {
		year          int
		expectedBadge string
	}{
		{year: 2021, expectedBadge: "A"},
		{year: 2020, expectedBadge: "A"},
		{year: 2018, expectedBadge: "B"},
		{year: 2015, expectedBadge: "B"},
		{year: 2012, expectedBadge: "C"},
		{year: 2010, expectedBadge: "C"},
		{year: 2007, expectedBadge: "D"},
		{year: 2005, expectedBadge: "D"},
		{year: 2000, expectedBadge: "E"},
	}

	for _, tt := range tests {
		t.Run(
			fmt.Sprintf("Year %d should return badge %s", tt.year, tt.expectedBadge),
			func(t *testing.T) {
				// Given
				car := vehicle.Vehicle{
					ID:    uuid.New(),
					Brand: "TestBrand",
					Model: "TestModel",
					Year:  tt.year,
				}

				// When
				badge := car.ContaminationBadge()

				// Then
				require.Equal(t, tt.expectedBadge, badge)
			},
		)
	}
}
