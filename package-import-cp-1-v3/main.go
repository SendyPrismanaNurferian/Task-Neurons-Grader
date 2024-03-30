package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Calculator struct{}

func (c Calculator) Add(a, b float32) float32 {
	return a + b
}

func (c Calculator) Subtract(a, b float32) float32 {
	return a - b
}

func (c Calculator) Multiply(a, b float32) float32 {
	return a * b
}

func (c Calculator) Divide(a, b float32) float32 {
	if b == 0 {
		return 0
	}
	return a / b
}

func (c Calculator) Result(input string) float32 {
	operands := strings.Split(input, " ")
	if len(operands) == 0 {
		return 0
	}

	var result float32
	temp, _ := strconv.ParseFloat(operands[0], 32)
	result = float32(temp)
	for i := 1; i < len(operands); i += 2 {
		operator := operands[i]
		operand, _ := strconv.ParseFloat(operands[i+1], 32)
		switch operator {
		case "+":
			result = c.Add(result, float32(operand))
		case "-":
			result = c.Subtract(result, float32(operand))
		case "*":
			result = c.Multiply(result, float32(operand))
		case "/":
			result = c.Divide(result, float32(operand))
		}
	}
	return result
}

func AdvanceCalculator(calculate string) float32 {
	calculator := Calculator{}
	return calculator.Result(calculate)
}

func main() {
	res := AdvanceCalculator("3 * 4 / 2 + 10 - 5")
	fmt.Println(res)

	res = AdvanceCalculator("10 / 4 + 100")
	fmt.Println(res)

	res = AdvanceCalculator("10 + 10 + 10 + 10 + 12 + 12 + 12 + 12")
	fmt.Println(res)

	res = AdvanceCalculator("10")
	fmt.Println(res)

	res = AdvanceCalculator("")
	fmt.Println(res)
}
