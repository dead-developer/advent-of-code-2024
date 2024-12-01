package main

import (
	"AoC2024/framework"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var list1 []int
var list2 []int

func main() {
	total := solution()
	fmt.Println("solution:", total)
}

func solution() int {
	lines := framework.ReadInput("input.txt")
	buildLists(lines)

	total := 0

	for i := 0; i < len(list1); i++ {
		distance := absInt(list1[i] - list2[i])
		total += distance
	}
	return total
}

func absInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func buildLists(lines []string) {
	for _, line := range lines {
		num1, num2 := splitRow(line)
		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}

	//sort lists
	sort.Ints(list1)
	sort.Ints(list2)
}

func splitRow(row string) (int, int) {
	parts := strings.Split(row, "  ")

	part1 := strings.TrimSpace(parts[0])
	num1, _ := strconv.Atoi(part1)

	part2 := strings.TrimSpace(parts[1])
	num2, _ := strconv.Atoi(part2)

	return num1, num2
}
