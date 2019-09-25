package main

import (
	"fmt"
	"os"
	"strings"
)

var opList = map[string]int{
	"*": 2,
	"/": 2,
	"+": 1,
	"-": 1,
	"^": 3,
}

func Parsing(input string) string {
	var operations, result string
	for i, _ := range input {
		switch string(input[i]) {
		case "(":
			operations += string(input[i])
		case ")":
			ind := strings.LastIndex(operations, "(")
			for i := len(operations); i > len(operations)-ind; i-- {
				result += string(operations[i]) + " "
			}
			operations = operations[:len(operations)-ind]
		case "*":
			operations += "*"
			for len(operations) > 0 {
				if op := opList[string(operations[len(operations)-1])]; op == 0 || op < opList["*"] {
					break
				}
				result += string(operations[len(operations) - 1]) + " "
				operations = operations[:len(operations)-1]
			}
			fmt.Println(result)
		case "/":
		case "+":
		case "-":
		default:
			result += string(input[i]) + " "
		}
	}
	return result
}

func getPriority() {

}

func main() {
	fmt.Print(Parsing(os.Args[1]))
}
