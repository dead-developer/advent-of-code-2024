package main

import (
	"AoC2024/framework"
	"fmt"
	"strings"
)

var matrix [][]string
var moves []string

var robotX, robotY int
var currentMove int

type vector struct {
	x int
	y int
}

func main() {
	total := solution()
	fmt.Println("solution:", total)
}

func solution() int {
	loadData("input.txt")

	maxTurns := len(moves)
	for currentMove < maxTurns {
		processMove()
		currentMove++
	}
	total := countBoxes()

	return total
}

func countBoxes() int {
	count := 0
	for y, line := range matrix {
		for x, char := range line {
			if char == "O" {
				count += 100*y + x
			}
		}
	}
	return count
}

func processMove() {
	move := moves[currentMove]

	var moveVector vector
	if move == "<" {
		moveVector = vector{-1, 0}
	} else if move == ">" {
		moveVector = vector{1, 0}
	} else if move == "^" {
		moveVector = vector{0, -1}
	} else if move == "v" {
		moveVector = vector{0, 1}
	}

	targetX, targetY := robotX+moveVector.x, robotY+moveVector.y

	if matrix[targetY][targetX] == "#" {
		return
	}
	if matrix[targetY][targetX] == "O" {
		pushO(targetX, targetY, moveVector)
	}
	if matrix[targetY][targetX] == "." {
		robotX = targetX
		robotY = targetY

	}

}

func pushO(x, y int, moveVector vector) {

	// try to push all O's
	targetX, targetY := x+moveVector.x, y+moveVector.y
	//fmt.Println("pushing", targetX, targetY, moveVector, targetX, targetY)
	if matrix[targetY][targetX] == "O" {
		pushO(targetX, targetY, moveVector)
	}
	if matrix[targetY][targetX] == "#" {
		return
	}
	if matrix[targetY][targetX] == "." {
		matrix[targetY][targetX] = "O"
		matrix[y][x] = "."
	}
}

func loadData(filename string) {
	lines := framework.ReadInput(filename)

	parse := "matrix"
	for _, line := range lines {
		if line == "" {
			parse = "moves"
		}
		if parse == "matrix" {
			matrix = append(matrix, strings.Split(line, ""))
		}
		if parse == "moves" {
			// split moves into array
			moves = append(moves, strings.Split(line, "")...)
		}
	}

	// find robot @
	for i, row := range matrix {
		for j, col := range row {
			if col == "@" {
				robotX = i
				robotY = j
				matrix[robotX][robotY] = "."
			}
		}
	}
}
