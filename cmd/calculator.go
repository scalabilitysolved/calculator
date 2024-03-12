package cmd

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var regex = regexp.MustCompile(`(-\d+\.\d+|-\d+|\d+\.\d+|\d+|[\+\-\*/\(\)])`)

type Calculator struct{}

type Operator int

type OperatorInfo struct {
	operator   Operator
	precedence int
}

func GetPrecedence(op Operator) int {
	return operatorPrecedence[op]
}

var operatorPrecedence = map[Operator]int{
	Multiplication: 1, // Treat Multiplication and Division equally
	Division:       1, // Equal to Multiplication
	Addition:       2, // Treat Addition and Subtraction equally
	Subtraction:    2, // Equal to Addition
}

func findOperator(operator string) Operator {
	o, _ := operatorToEnum[operator]
	return o
}

func NewOperatorInfo(value string) OperatorInfo {
	o := findOperator(value)
	p := GetPrecedence(o)
	return OperatorInfo{
		operator:   o,
		precedence: p,
	}
}

const (
	Addition Operator = iota
	Subtraction
	Multiplication
	Division
)

var operatorToEnum = map[string]Operator{
	"+": Addition,
	"-": Subtraction,
	"*": Multiplication,
	"/": Division,
}

func (c Calculator) evaluateExpression(parts []string) (float64, error) {
	if len(parts) < 3 || len(parts)%2 == 0 {
		return 0, fmt.Errorf("invalid expression format")
	}

	var numbersStack []float64
	var operatorsStack []OperatorInfo

	applyOperation := func() error {
		if len(numbersStack) < 2 {
			return fmt.Errorf("not enough numbers to apply operation")
		}
		b := numbersStack[len(numbersStack)-1]
		numbersStack = numbersStack[:len(numbersStack)-1]

		a := numbersStack[len(numbersStack)-1]
		numbersStack = numbersStack[:len(numbersStack)-1]

		op := operatorsStack[len(operatorsStack)-1]
		operatorsStack = operatorsStack[:len(operatorsStack)-1]

		var result float64
		switch op.operator {
		case Addition:
			result = a + b
		case Subtraction:
			result = a - b
		case Multiplication:
			result = a * b
		case Division:
			if b == 0 {
				return fmt.Errorf("division by zero")
			}
			result = a / b
		default:
			return fmt.Errorf("unsupported operation")
		}
		numbersStack = append(numbersStack, result)
		return nil
	}

	for i, part := range parts {
		if i%2 == 0 { // number
			num, err := strconv.ParseFloat(part, 64)
			if err != nil {
				return 0, fmt.Errorf("couldn't convert %s to number", part)
			}
			numbersStack = append(numbersStack, num)
		} else { // operator
			currentOp := NewOperatorInfo(part)
			for len(operatorsStack) > 0 && operatorsStack[len(operatorsStack)-1].precedence <= currentOp.precedence {
				if err := applyOperation(); err != nil {
					return 0, err
				}
			}
			operatorsStack = append(operatorsStack, currentOp)
		}
	}

	for len(operatorsStack) > 0 {
		if err := applyOperation(); err != nil {
			return 0, err
		}
	}

	if len(numbersStack) != 1 {
		return 0, fmt.Errorf("error in calculation, final numbers stack should have exactly one element")
	}

	return numbersStack[0], nil
}

func (c Calculator) calculate(input string) (float64, error) {
	if len(input) == 0 {
		return 0, fmt.Errorf("input cannot be empty")
	}

	evaluate := func(expression string) (float64, error) {
		// This inner function is responsible for recursively evaluating expressions within parentheses.
		// It should call `evaluateExpression` for expressions without any parentheses.
		parts := regex.FindAllString(expression, -1)
		return c.evaluateExpression(parts)
	}

	var resolveParentheses func(expression string) (string, error)
	resolveParentheses = func(expression string) (string, error) {
		// Recursively resolve all parentheses in the expression
		for {
			left := strings.LastIndex(expression, "(")
			if left == -1 {
				break
			}
			right := strings.Index(expression[left:], ")") + left
			if right < left {
				return "", fmt.Errorf("mismatched parentheses")
			}

			// Extract the expression within the parentheses
			innerExpression := expression[left+1 : right]
			result, err := evaluate(innerExpression)
			if err != nil {
				return "", err
			}
			// Replace the parenthesis expression with its result
			expression = expression[:left] + fmt.Sprintf("%f", result) + expression[right+1:]
		}
		return expression, nil
	}

	resolvedExpression, err := resolveParentheses(input)
	if err != nil {
		return 0, err
	}

	return evaluate(resolvedExpression)
}
