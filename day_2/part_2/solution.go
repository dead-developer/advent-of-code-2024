package main

import (
	"AoC2024/framework"
	"fmt"
	"strconv"
	"strings"
)

var numbers [][]int

func main() {
	total := solution()
	fmt.Println("solution:", total)
}

func solution() int {
	loadData()

	total := 0
	for _, line := range numbers {
		if isSafe(line) {
			total++
			continue
		}
		// check if is safe when removing single number
		for i := range line {
			newLine := make([]int, len(line)-1)
			copy(newLine, line[:i])
			copy(newLine[i:], line[i+1:])
			if isSafe(newLine) {
				total++
				break
			}
		}

	}

	return total
}

func isSafe(line []int) bool {
	// check if the line is safe
	safe := true
	increase := false
	decrease := false

	for i := range line {
		if i == 0 {
			continue
		}
		if line[i] > line[i-1] {
			increase = true
		} else if line[i] < line[i-1] {
			decrease = true
		} else if line[i] == line[i-1] {
			safe = false
		}
		if increase && decrease {
			safe = false
		}

		// if increase/decrease more than 3
		difference := line[i] - line[i-1]
		if difference > 3 || difference < -3 {
			safe = false
		}

	}
	return safe

}

func loadData() {
	lines := framework.ReadInput("input.txt")
	// split each line by spaces
	for _, line := range lines {
		numbers = append(numbers, splitLine(line))
	}

}

func splitLine(line string) []int {
	parts := strings.Fields(line)
	var result []int
	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			fmt.Println("Error converting string to int:", err)
			continue
		}
		result = append(result, num)
	}
	return result
}
