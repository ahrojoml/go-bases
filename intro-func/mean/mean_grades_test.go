package main_test

import (
	mean "go-bases/intro-func/mean"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetMean(t *testing.T) {
	grades := []int{2, 4, 5, 6, 8, 10}
	expect := 5

	result := mean.GetMean(grades...)

	require.Equal(t, expect, result)

}
