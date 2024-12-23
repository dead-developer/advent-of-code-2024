package main

import (
	"AoC2024/framework"
	"fmt"
	"strconv"
	"strings"
)

var testValues []int
var equations [][]int

func main() {
	total := solution()
	fmt.Println("solution:", total)
}

func solution() int {
	loadData()

	results := make(chan int, len(equations))

	total := 0
	var operators1 = []string{"+", "*"}
	var operators2 = []string{"+", "*", "||"}

	for i, equation := range equations {
		go func(eq []int, tv int, i int) {
			result := runTests(eq, tv, operators1)
			if result == 0 {
				result += runTests(eq, tv, operators2)
			}
			results <- result
		}(equation, testValues[i], i)
	}

	for range equations {
		total += <-results
	}

	return total
}

func runTests(equation []int, testValue int, checkOperators []string) int {
	if len(equation) == 1 {
		if equation[0] == testValue {
			return equation[0]
		}
	}

	operatorSlots := len(equation) - 1
	totalCombinations := intPow(len(checkOperators), operatorSlots)

	for i := 0; i < totalCombinations; i++ {
		ops := make([]string, operatorSlots)
		temp := i

		// get combinations v2
		amountOfOperators := len(checkOperators)
		for j := 0; j < operatorSlots; j++ {
			ops[j] = checkOperators[temp%amountOfOperators]
			temp /= amountOfOperators
		}

		if calculateEquation(equation, ops) == testValue {
			return testValue
		}
	}

	return 0
}

func calculateEquation(equation []int, operators []string) int {
	result := equation[0]
	for i := 0; i < len(operators); i++ {
		if operators[i] == "+" {
			result += equation[i+1]
		} else if operators[i] == "*" {
			result *= equation[i+1]
		} else if operators[i] == "||" {
			result = concatenateNumbers(result, equation[i+1])
		}
	}

	return result
}

func concatenateNumbers(number1 int, number2 int) int {
	//convert to string
	str1 := strconv.Itoa(number1)
	str2 := strconv.Itoa(number2)
	value := str1 + str2

	//convert to int
	output, _ := strconv.Atoi(value)

	return output
}

func intPow(base, exponent int) int {
	result := 1
	for i := 0; i < exponent; i++ {
		result *= base
	}
	return result
}

func loadData() {
	lines := framework.ReadInput("input.txt")

	// split to matrix
	for _, line := range lines {
		//split line by :
		parts := strings.Split(line, ": ")
		testValue, _ := strconv.Atoi(parts[0])
		testValues = append(testValues, testValue)

		// split other numbers by space
		numbers := strings.Split(parts[1], " ")
		equation := make([]int, 0)
		for _, number := range numbers {
			num, _ := strconv.Atoi(number)
			equation = append(equation, num)
		}
		equations = append(equations, equation)
	}

}
