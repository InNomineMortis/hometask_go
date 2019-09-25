package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var opList = map[string]int{
	"*": 2,
	"/": 2,
	"+": 1,
	"-": 1,
	"^": 3,
}

func operator(op string) bool {
	if op == "*" || op == "/" || op == "+" || op == "-" || op == "^" {
		return true
	}
	return false
}

func Parsing(input string) string {
	var operations, result string
	flag := false
	for i, _ := range input {
		switch string(input[i]) {
		case "(":
			operations += string(input[i])
			flag = false
		case ")":
			ind := strings.LastIndex(operations, "(")
			for i := len(operations) - 1; i > ind; i-- {
				result += " " + string(operations[i])
			}
			operations = operations[:ind]
			flag = false
		default:
			if operator(string(input[i])) {
				for len(operations) > 0 {
					op := opList[string(operations[len(operations)-1])]
					if op == 0 || op < opList[string(input[i])] {
						break
					}
					result += " " + string(operations[len(operations)-1])
					operations = operations[:len(operations)-1]
				}
				operations += string(input[i])
				flag = false
			} else {
				if flag {
					result += string(input[i])
				} else {
					result += " " + string(input[i])
				}
				flag = true
			}
		}
	}
	for len(operations) > 0 {
		result += " " + string(operations[len(operations)-1])
		operations = operations[:len(operations)-1]
	}
	result = result[1:]
	return result
}

func calculate(rpn string) float64 {
	res := strings.Split(rpn, " ")
	stack := make([]int, 1)
	for _, v := range res {
		num, err := strconv.Atoi(v)
		if err != nil {
			switch v {
			case " ":
				break
			default:
				if operator(v) {
					stack = equation(v, stack)
				} else {
					fmt.Println("Unknown operation!:", v, ":")
				}
			}
		} else {
			stack = append(stack, num)
		}
	}
	finRes, _ := pop(stack)
	return float64(finRes)
}
func equation(op string, stack []int) []int {
	x, stack := pop(stack)
	y, stack := pop(stack)
	pop(stack)
	switch op {
	case "*":
		return append(stack, x*y)
	case "/":
		if y == 0 {
			panic("Division zero")
		}
		return append(stack, x/y)
	case "+":
		return append(stack, x+y)
	case "-":
		return append(stack, x-y)
	case "^":
		return append(stack, x^y)
	}
	return nil
}

func pop(stack []int) (int, []int) {
	return stack[len(stack)-1], stack[:len(stack)-1]
}
func main() {
	res := Parsing(os.Args[1])
	fmt.Println(calculate(res))
}
