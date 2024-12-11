package main

import (
	"AoC2024/framework"
	"fmt"
	"math"
	"strconv"
	"strings"
)

var stones []int

func main() {
	total := solution()
	fmt.Println("solution:", total)
}

func solution() int {
	loadData()

	blinks := 25

	for i := 0; i < blinks; i++ {
		stones = blink()
	}

	return len(stones)
}

func blink() []int {
	output := make([]int, 0)

	for i := 0; i < len(stones); i++ {

		if stones[i] == 0 {
			stones[i] = 1
			output = append(output, stones[i])
		} else if hasEvenDigits(stones[i]) {
			current := strconv.Itoa(stones[i])
			half := len(current) / 2
			firstPart, _ := strconv.Atoi(current[:half])
			lastPart, _ := strconv.Atoi(current[half:])
			output = append(output, firstPart)
			output = append(output, lastPart)

		} else {

			output = append(output, stones[i]*2024)
		}
	}
	return output
}

func hasEvenDigits(n int) bool {
	if n == 0 {
		return false
	}

	if n < 0 {
		n = -n
	}

	// Use integer log10 to count digits
	return (int(math.Log10(float64(n)))+1)%2 == 0
}

func removeLeading(s string) string {
	s = strings.TrimLeft(s, "0")
	if s == "" {
		return "0"
	}
	return s
}

func loadData() {
	lines := framework.ReadInput("input.txt")

	for _, digits := range strings.Split(lines[0], " ") {
		number, _ := strconv.Atoi(digits)
		stones = append(stones, number)
	}
}
