package main_test

import (
	salary "go-bases/intro-func/salary"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetTax_Below50(t *testing.T) {
	payment := 49999.0
	expect := 0.0

	result, err := salary.GetTax(payment)

	require.Equal(t, expect, result)
	require.NoError(t, err)
}

func TestGetTax_Over50Below150(t *testing.T) {
	payment := 140000.0
	expect := payment * 0.1

	result, err := salary.GetTax(payment)

	require.Equal(t, expect, result)
	require.NoError(t, err)
}

func TestGetTax_Over150(t *testing.T) {
	payment := 170000.0
	var expect float64
	expect += payment * 0.17
	expect += payment * 0.1

	result, err := salary.GetTax(payment)

	diff := math.Abs(result - expect)

	require.LessOrEqual(t, diff, 1e-9) // epsilon to compare floats
	require.NoError(t, err)
}
