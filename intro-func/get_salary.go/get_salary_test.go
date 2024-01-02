package getsalary_test

import (
	getsalary "go-bases/intro-func/get_salary.go"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSalaryTypeA(t *testing.T) {
	minutes := 120
	expect := 9000.0

	result := getsalary.SalaryTypeA(minutes)

	require.Equal(t, expect, result)
}

func TestSalaryTypeB(t *testing.T) {
	minutes := 120
	expect := 3600.0

	result := getsalary.SalaryTypeB(minutes)

	require.Equal(t, expect, result)
}

func TestSalaryTypeC(t *testing.T) {
	minutes := 120
	expect := 2000.0

	result := getsalary.SalaryTypeC(minutes)

	require.Equal(t, expect, result)
}
