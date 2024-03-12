package cmd

import (
	"fmt"
	"math"
	"testing"
)

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
			_, err := Calculator{}.calculate(test)

			if err == nil {
				t.Errorf("Expected invalid input state for %s", test)
			}

		})
	}
}

func TestAddition(t *testing.T) {

	tests := []struct {
		input    string
		expected float64
	}{
		{"4 + 2", 6},
		{"3 + 2 + 1", 6},

		{"3 + 0", 3},
		{"1000 + 10 + 100 + 90", 1200},
		{"2 + -2", 0},
	}

	executeTests(t, tests)
}

func TestSubtraction(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
	}{
		{"8 - 2", 6},
		{"10 + 2 - 1", 11},
		{"2 + -2 - -10 + 2", 12},
		{"2 + -10 + 2", -6},
		{"5 - -10", 15},
	}

	executeTests(t, tests)
}

func TestMultiplication(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
	}{
		{"5 * 5", 25},
		{"5 * 10", 50},
		{"100 * 10", 1000},
		{"2 * 10 + 5", 25},
	}

	executeTests(t, tests)
}

func TestDivision(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
	}{
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
	tests := []struct {
		input    string
		expected float64
	}{
		{"(10 + 10) / 5", 4},
		{"(10 + 10) * 5", 100},
		{"(10 + 10) + (10 * 2)", 40},
		{"6 + (10 / 1.9) * (10 * 2)", 111.263157895},
		{"(3 + 2) + (8 + (9 * 2))", 31},
		{"(103 * (3 / 2) + 9) + (8 + (9 * 2))", 189.5},
	}
	executeTests(t, tests)
}

func executeTests(t *testing.T, tests []struct {
	input    string
	expected float64
}) {

	const tolerance = 0.00001
	for _, test := range tests {
		testName := fmt.Sprintf("%s equals %f", test.input, test.expected)
		t.Run(testName, func(t *testing.T) {
			result, err := Calculator{}.calculate(test.input)

			if err != nil {
				t.Errorf("Encountered error %v", err)
			}

			diff := math.Abs(result - test.expected)
			if diff > tolerance {
				t.Errorf("got %f, want %f, difference %f exceeds tolerance %f", result, test.expected, diff, tolerance)
			}
		})
	}
}
