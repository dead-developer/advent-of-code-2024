package main

import (
	"AoC2024/framework"
	"fmt"
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
			if char == "X" {
				total += findWords(x, y)
			}
		}
	}
	return total
}

func findWords(x, y int) int {
	counts := 0
	counts += findLetters(x, y, 1, 0)
	counts += findLetters(x, y, -1, 0)
	counts += findLetters(x, y, 0, 1)
	counts += findLetters(x, y, 0, -1)

	counts += findLetters(x, y, -1, -1)
	counts += findLetters(x, y, -1, 1)
	counts += findLetters(x, y, 1, -1)
	counts += findLetters(x, y, 1, 1)

	return counts
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
