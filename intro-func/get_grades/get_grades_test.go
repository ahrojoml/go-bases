package getgrades_test

import (
	getgrades "go-bases/intro-func/get_grades"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetMean(t *testing.T) {
	grades := []int{2, 4, 5, 6, 8, 10}
	expect := 5

	result, err := getgrades.GetMean(grades...)

	require.Equal(t, expect, result)
	require.NoError(t, err)
}

func TestGetMean_Empty(t *testing.T) {
	grades := []int{}

	_, err := getgrades.GetMean(grades...)

	require.Error(t, err)
}

func TestGetMin(t *testing.T) {
	grades := []int{2, 4, 5, 6, 8, 10}
	expect := 2

	result, err := getgrades.GetMin(grades...)

	require.Equal(t, expect, result)
	require.NoError(t, err)
}

func TestGetMin_Empty(t *testing.T) {
	grades := []int{}

	_, err := getgrades.GetMin(grades...)

	require.Error(t, err)
}

func TestGetMax(t *testing.T) {
	grades := []int{2, 4, 5, 6, 8, 10}
	expect := 10

	result, err := getgrades.GetMax(grades...)

	require.Equal(t, expect, result)
	require.NoError(t, err)
}

func TestGetMax_Empty(t *testing.T) {
	grades := []int{}

	_, err := getgrades.GetMax(grades...)

	require.Error(t, err)
}
