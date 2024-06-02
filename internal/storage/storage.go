package storage

import (
	"GOncurrently-Calculator/internal/structures/stack"
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"
)

type Expression struct {
	Id         int    `json:"id"`
	Expression string `json:"expression"`
	Status     string `json:"status"`
	Result     string `json:"result"`
	rpn        string
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
	if e.Expression == "" {
		return true
	}

	stack := []rune{}

	for i := 0; i < len(e.Expression); i++ {
		char := e.Expression[i]
		if !(unicode.IsDigit(rune(char)) || char == '+' || char == '-' || char == '*' || char == '/' || char == '(' || char == ')' || char == '^') {
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

func (e *Expression) AddSpaces() {
	res := ""
	for i, v := range e.Expression {
		if unicode.IsDigit(v) {
			res += string(v)
		} else {
			idx := i + 1
			if i == len(e.Expression)-1 {
				idx = i
			}
			if unicode.IsDigit(rune(e.Expression[idx])) {
				res += " " + string(v) + " "
			} else {
				res += " " + string(v)
			}
		}
	}

	e.Expression = strings.Trim(res, " ")
}

func isStringDigit(s string) bool {
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

func (e *Expression) ToRpn() string { // RPN - Reverse Polish Notation
	priorities := make(map[string]int)
	priorities["^"] = 4
	priorities["*"] = 3
	priorities["/"] = 3
	priorities["+"] = 2
	priorities["-"] = 2
	priorities["("] = 1

	var result string = ""
	stk := stack.New[string]()

	for _, v := range strings.Split(e.Expression, " ") {
		if isStringDigit(v) {
			result += v + " "
		} else {
			if v == "(" {
				stk.Push(v)
				continue
			} else if v == ")" {
				for {
					if stk.GetTop() == "(" {
						stk.Pop()
						break
					}
					val, _ := stk.Pop()
					result += val + " "
				}
				continue
			}

			if stk.IsEmpty() || priorities[stk.GetTop()] < priorities[v] {
				stk.Push(v)
			} else {
				for val, is := stk.Pop(); priorities[val] >= priorities[v] && is; {
					if val == "(" {
						val, is = stk.Pop()
						continue
					}
					result += val + " "
					val, is = stk.Pop()
				}
				stk.Push(v)
			}
		}
	}

	for val, is := stk.Pop(); is; {
		if val == ")" {
			val, is = stk.Pop()
			continue
		}

		result += val + " "
		val, is = stk.Pop()
	}

	e.rpn = strings.Trim(result, " ")

	return result
}

func (e *Expression) Calculate() (int, error) {
	stk := stack.New[int]()

	for _, v := range strings.Split(e.rpn, " ") {
		if isStringDigit(v) {
			digit, err := strconv.Atoi(v)
			if err != nil {
				return 0, fmt.Errorf("invalid expression")
			}

			stk.Push(digit)
		} else {
			a, ok := stk.Pop()
			if !ok {
				fmt.Printf("here: %v\n", v)
				return 0, fmt.Errorf("invalid expression")
			}
			b, ok := stk.Pop()
			if !ok {
				fmt.Printf("here: %v\n", v)
				return 0, fmt.Errorf("invalid expression")
			}

			if v == "+" {
				stk.Push(b + a)
			} else if v == "-" {
				stk.Push(b - a)
			} else if v == "*" {
				stk.Push(b * a)
			} else if v == "^" {
				stk.Push(int(math.Pow(float64(b), float64(a))))
			} else if v == "/" {
				stk.Push(b / a)
			}
		}
	}

	return stk.GetTop(), nil
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
