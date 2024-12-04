package main

import (
	"AoC2024/framework"
	"fmt"
	"sort"
	"strings"
)

var matrix [][]string

func main() {
	total := solution()
	fmt.Println("solution:", total)
}

func solution() int {
	loadData()

	total := 0
	//find 'X'
	for y, row := range matrix {
		for x, char := range row {
			if (x < 1) || (y < 1) || (y >= len(matrix)-1) || (x >= len(matrix[0])-1) {
				continue
			}
			if char == "A" {
				total += findWords(x, y)
			}
		}
	}

	return total
}

func findWords(x, y int) int {
	// check for diagonal
	arr1 := []string{matrix[y-1][x-1], matrix[y+1][x+1]}
	arr2 := []string{matrix[y-1][x+1], matrix[y+1][x-1]}

	sort.Strings(arr1)
	sort.Strings(arr2)

	if (arr2[0] == "M" && arr2[1] == "S") && (arr1[0] == "M" && arr1[1] == "S") {
		return 1
	}

	return 0
}

func findLetters(startX, startY int, directionX, directionY int) int {
	letters := []string{"M", "A", "S"}
	// check if letters are found in matrix in direction
	for _, letter := range letters {
		startX += directionX
		startY += directionY
		if (startX < 0) || (startY < 0) || (startX >= len(matrix)) || (startY >= len(matrix[0])) {
			return 0
		}
		if matrix[startY][startX] != letter {
			return 0
		}
	}
	return 1
}

func loadData() {
	lines := framework.ReadInput("input.txt")

	// split to matrix
	for _, line := range lines {
		matrix = append(matrix, strings.Split(line, ""))
	}
}
