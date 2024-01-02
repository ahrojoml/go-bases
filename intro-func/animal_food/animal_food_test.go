package animalfood_test

import (
	animalfood "go-bases/intro-func/animal_food"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDog(t *testing.T) {
	need := 30.0
	expect := 20.0

	result := animalfood.DogFood(need)

	require.Equal(t, expect, result)
}

func TestCat(t *testing.T) {
	need := 30.0
	expect := 25.0

	result := animalfood.CatFood(need)

	require.Equal(t, expect, result)
}

func TestHamster(t *testing.T) {
	need := 30.0
	expect := 29.750

	result := animalfood.HamsterFood(need)

	require.Equal(t, expect, result)
}

func TestTarantula(t *testing.T) {
	need := 30.0
	expect := 29.850

	result := animalfood.TarantulaFood(need)

	require.Equal(t, expect, result)
}
