package calc

type Operator int

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

type OperatorInfo struct {
	operator   Operator
	precedence int
}

func getPrecedence(op Operator) int {
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
	p := getPrecedence(o)
	return OperatorInfo{
		operator:   o,
		precedence: p,
	}
}
