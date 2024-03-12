package calc

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAddition(t *testing.T) {

	tests := []testCase{
		{"4 + 2", 6},
		{"3 + 2 + 1", 6},
		{"3 + 0", 3},
		{"1000 + 10 + 100 + 90", 1200},
		{"2 + -2", 0},
	}

	executeTests(t, tests)
}

func TestSubtraction(t *testing.T) {
	tests := []testCase{
		{"8 - 2", 6},
		{"10 + 2 - 1", 11},
		{"2 + -2 - -10 + 2", 12},
		{"2 + -10 + 2", -6},
		{"5 - -10", 15},
	}

	executeTests(t, tests)
}

func TestMultiplication(t *testing.T) {
	tests := []testCase{
		{"5 * 5", 25},
		{"5 * 10", 50},
		{"100 * 10", 1000},
		{"2 * 10 + 5", 25},
	}

	executeTests(t, tests)
}

func TestDivision(t *testing.T) {
	tests := []testCase{
		{"50 / 5", 10},
		{"10 / 10", 1},
		{"100 / 10", 10},
		{"10 / 2 + 5", 10},
		{"2 + 10 / 5", 4},
		{" 5 / 10", 0.5},
		{"100 * 10 + 9 / 34", 1000.26470588},
		{"90 + 30 / 8 * 105.4", 485.25},
	}
	executeTests(t, tests)
}

func TestOperatorPrecedence(t *testing.T) {
	tests := []testCase{
		{"(10 + 10) / 5", 4},
		{"(10 + 10) * 5", 100},
		{"(10 + 10) + (10 * 2)", 40},
		{"6 + (10 / 1.9) * (10 * 2)", 111.263157895},
		{"(3 + 2) + (8 + (9 * 2))", 31},
		{"(103 * (3 / 2) + 9) + (8 + (9 * 2))", 189.5},
	}
	executeTests(t, tests)
}

func TestExpectedErrors(t *testing.T) {
	tests := []string{
		"2 + a",
		"2 + a + b / c",
		"",
		"3 +",
		"", "9 * ",
		"9 *** 9",
	}
	for _, test := range tests {
		testName := fmt.Sprintf("%s produces error", test)
		t.Run(testName, func(t *testing.T) {
			_, err := Calculate(test)
			require.Errorf(t, err, "Expected invalid input state for %s", test)
		})
	}
}

type testCase struct {
	input    string
	expected float64
}

const delta = 0.00001

func executeTests(t *testing.T, tests []testCase) {
	t.Helper()
	for _, test := range tests {
		t.Run(fmt.Sprintf("%s = %f", test.input, test.expected), func(t *testing.T) {
			result, err := Calculate(test.input)
			require.NoError(t, err)
			assert.InDelta(t, test.expected, result, delta)
		})
	}
}
