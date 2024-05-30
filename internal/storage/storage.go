package storage

import "unicode"

type Expression struct {
	Id         int    `json:"id"`
	Expression string `json:"expression"`
	Status     string `json:"status"`
	Result     string `json:"result"`
}

func (e *Expression) RemoveSpaces() *Expression {
	result := ""
	for i := 0; i < len(e.Expression); i++ {
		if e.Expression[i] != ' ' {
			result += string(e.Expression[i])
		}
	}
	e.Expression = result
	return e
}

func (e *Expression) IsInvalid() bool {
	stack := []rune{}

	for i := 0; i < len(e.Expression); i++ {
		char := e.Expression[i]
		if !(unicode.IsDigit(rune(char)) || char == '+' || char == '-' || char == '*' || char == '/' || char == '(' || char == ')') {
			return true
		}

		if char == '(' {
			stack = append(stack, rune(char))
		}

		if char == ')' {
			if len(stack) == 0 {
				return true
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) != 0
}

type Expressions struct {
	Storage []*Expression `json:"expressions"`
}

var (
	DB *Expressions
)

func init() {
	DB = &Expressions{
		Storage: make([]*Expression, 0),
	}
}
