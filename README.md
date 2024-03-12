# Go calculator

## Description

This is a simple Reverse Polish Notation (RPN) calculator written in Go. It can handle basic arithmetic operations including 
addition, subtraction, multiplication, and division. The calculator also supports expressions with parentheses, 
allowing for complex calculations.

### Features
 - Supports basic arithmetic operations: +, -, *, /
 - Handles expressions with parentheses (including nested parentheses)
 - Uses Reverse Polish Notation (RPN) for evaluation

### Usage

```Go
package main

import (
	"./cmd" // Replace with the actual path to the package containing the Calculator struct
	"fmt"
)

func main() {
	calculator := cmd.Calculator{}
	result, err := calculator.calculate("2 * (3 + 1) / 4")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Result:", result)
}

```


How it Works
The calculator parses the input expression into tokens using a regular expression. It then uses a shunting-yard algorithm to convert the infix expression (standard mathematical notation) into Reverse Polish Notation (RPN).

RPN places operands (numbers) before operators, allowing for evaluation without the need for explicit operator precedence rules.

The calculator then evaluates the RPN expression using a stack-based approach. It iterates through the tokens and pushes operands onto a stack. When an operator is encountered, it pops the required number of operands from the stack, performs the operation, and pushes the result back onto the stack.
Finally, the calculator checks if the stack has exactly one element remaining. This element represents the final result of the expression.

### Downsides
* It uses floats so, yeah 🙃
* Doesn't handle exponents or other operators