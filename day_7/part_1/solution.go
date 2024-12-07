package main

import (
	"AoC2024/framework"
	"fmt"
	"strconv"
	"strings"
)

var testValues []int
var equations [][]int

var operators = []string{"+", "*"}

func main() {
	total := solution()
	fmt.Println("solution:", total)
}

func solution() int {
	loadData()

	total := 0

	for i, equation := range equations {
		total += runTests(equation, testValues[i])
	}

	return total
}

func runTests(equation []int, testValue int) int {
	// run combinations of operators to numbers
	combinations := generateCombinations(operators, len(equation)-1, []string{})

	// test each combination
	for _, combination := range combinations {
		// do the math according to operators
		result := equation[0]
		for i, op := range combination {
			if op == "+" {
				result += equation[i+1]
			} else if op == "*" {
				result *= equation[i+1]
			}
		}
		if result == testValue {
			return result
		}
	}
	return 0
}

func generateCombinations(operators []string, slots int, current []string) [][]string {
	var result [][]string
	generateHelper(operators, slots, []string{}, &result)
	return result
}
func generateHelper(operators []string, slots int, current []string, result *[][]string) {
	if len(current) == slots {
		*result = append(*result, append([]string(nil), current...))
		return
	}

	for _, op := range operators {
		generateHelper(operators, slots, append(current, op), result)
	}
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
