package main

import (
	"AoC2024/framework"
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Solved with the help of Claude

func main() {
	total := solution()
	fmt.Println("solution:", total)
}

func solution() int {
	inputData := loadData("input.txt")

	total := calculateStones(25, inputData)

	return total
}

func calculateStones(blinks int, initialStones []int) int {
	// track stone types
	stoneCounts := make(map[int]int)
	for _, stone := range initialStones {
		stoneCounts[stone]++
	}

	// Process each step
	for i := 0; i < blinks; i++ {
		newCounts := make(map[int]int)

		// Process each type of stone
		for stone, count := range stoneCounts {
			if stone == 0 {
				newCounts[1] += count
			} else if hasEvenDigits(stone) {
				leftNumber, rightNumber := splitNumber(stone)
				newCounts[leftNumber] += count
				newCounts[rightNumber] += count
			} else {
				newCounts[stone*2024] += count
			}
		}
		stoneCounts = newCounts
	}

	// Count total stones
	total := 0
	for _, count := range stoneCounts {
		total += count
	}

	return total
}
func countDigits(n int) int {
	if n == 0 {
		return 1
	}
	return int(math.Log10(float64(n))) + 1
}

func hasEvenDigits(n int) bool {
	return countDigits(n)%2 == 0
}

func splitNumber(n int) (int, int) {
	str := fmt.Sprintf("%d", n)
	half := len(str) / 2
	left, _ := strconv.Atoi(str[:half])
	right, _ := strconv.Atoi(str[half:])
	return left, right
}

func loadData(filename string) []int {
	lines := framework.ReadInput(filename)
	var numbers []int

	for _, digits := range strings.Split(lines[0], " ") {
		number, _ := strconv.Atoi(digits)
		numbers = append(numbers, number)
	}

	return numbers
}
