package salary_test

import (
	"go-bases/intro-func/salary"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetTax_Below50(t *testing.T) {
	payment := 49999.0
	expect := 0.0

	result := salary.GetTax(payment)

	require.Equal(t, expect, result)
}

func TestGetTax_Over50Below150(t *testing.T) {
	payment := 140000.0
	expect := 14000.0

	result := salary.GetTax(payment)

	require.Equal(t, expect, result)
}

func TestGetTax_Over150(t *testing.T) {
	payment := 170000.0
	expect := 45900.0

	result := salary.GetTax(payment)

	require.Equal(t, expect, result)
}
