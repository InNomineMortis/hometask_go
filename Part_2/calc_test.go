package main

import (
	"testing"
)

func TestParsing(t *testing.T) {
	input := "(13-10+5)*10/(22+5)^2"
	res := Parsing(input)
	if res != "13 10 - 5 + 10 * 22 5 + 2 ^ /" {
		t.Error("Wrong result! Expected: 13 10 - 5 + 10 * 22 5 + 2 ^ / got: ", res)
	}
	input = "(10+15)*(8+7)/12+13^(3+4)"
	res = Parsing(input)
	if res != "10 15 + 8 7 + * 12 / 13 3 4 + ^ +" {
		t.Error("Wrong result! Expected: 10 15 + 8 7 + * 12 / 13 3 4 + ^ + ^ / got: ", res)
	}
	input = "13^2+11-(5-3)*16"
	res = Parsing(input)
	if res != "13 2 ^ 11 + 5 3 - 16 * -" {
		t.Error("Wrong result! Expected: 13 2 ^ 11 + 5 3 - 16 * - got: ", res)
	}
	input = "19-(16*2+3/1-(6+8^2)^3)"
	res = Parsing(input)
	if res != "19 16 2 * 3 1 / + 6 8 2 ^ + 3 ^ - -" {
		t.Error("Wrong result! Expected: 19 16 2 * 3 1 / + 6 8 2 ^ + 3 ^ - - got: ", res)
	}
}

func TestCalculate(t *testing.T)  {
	rpn := "248 10 - 5 + 3 * 22 5 + 2 ^ /"
	res := calculate(rpn)
	if res !=  1{
		t.Error("Wrong result! Expected: 1 Got: ", res)
	}
	rpn = "12 3 * 1 + 3 1 33 + * +"
	res = calculate(rpn)
	if res !=  139{
		t.Error("Wrong result! Expected: 139 Got: ", res)
	}
	rpn = "19 16 2 * 3 1 / + 6 8 2 ^ + 3 ^ - -"
	res = calculate(rpn)
	if res !=  342984{
		t.Error("Wrong result! Expected: 342984 Got: ", res)
	}
}