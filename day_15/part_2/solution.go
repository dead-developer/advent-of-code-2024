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

type coord struct {
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
			if char == "[" {
				count += 100*y + x
			}
		}
	}
	return count
}

func processMove() {
	move := moves[currentMove]

	var moveVector coord
	if move == "<" {
		moveVector = coord{-1, 0}
	} else if move == ">" {
		moveVector = coord{1, 0}
	} else if move == "^" {
		moveVector = coord{0, -1}
	} else if move == "v" {
		moveVector = coord{0, 1}
	}

	targetX, targetY := robotX+moveVector.x, robotY+moveVector.y

	if matrix[targetY][targetX] == "#" {

		return
	}
	if matrix[targetY][targetX] == "[" || matrix[targetY][targetX] == "]" {
		pushBox(targetX, targetY, moveVector)
	}
	if matrix[targetY][targetX] == "." {
		robotX = targetX
		robotY = targetY
	}

}

func pushBox(x, y int, moveVector coord) bool {
	if !canBoxMove(x, y, moveVector) {
		return false
	}

	targetCoords := getBoxTargetCoords(x, y, moveVector)

	for _, target := range targetCoords {
		if matrix[target.y][target.x] == "[" || matrix[target.y][target.x] == "]" {
			pushBox(target.x, target.y, moveVector)
		}
		if matrix[target.y][target.x] == "#" {
			return false
		}
	}
	// if space to move then move
	for _, target := range targetCoords {
		if matrix[target.y][target.x] != "." {
			return true
		}
	}
	moveBox(x, y, moveVector)
	return true
}

func findBox(x, y int) (coord, coord) {
	var leftSide, rightSide coord

	if matrix[y][x] == "]" {
		rightSide = coord{x, y}
		leftSide = coord{x - 1, y}
	} else {
		leftSide = coord{x, y}
		rightSide = coord{x + 1, y}
	}
	return leftSide, rightSide
}

func getBoxTargetCoords(x, y int, moveVector coord) []coord {
	var targetCoords []coord

	leftSide, rightSide := findBox(x, y)
	if moveVector.x == 0 {
		targetCoords = append(targetCoords, coord{rightSide.x, rightSide.y + moveVector.y})
		targetCoords = append(targetCoords, coord{leftSide.x, leftSide.y + moveVector.y})
	}
	if moveVector.y == 0 {
		if moveVector.x > 0 {
			targetCoords = append(targetCoords, coord{rightSide.x + moveVector.x, rightSide.y})
		} else {
			targetCoords = append(targetCoords, coord{leftSide.x + moveVector.x, leftSide.y})
		}
	}
	return targetCoords
}

func canBoxMove(x, y int, moveVector coord) bool {
	canMove := make([]bool, 0)
	targetCoords := getBoxTargetCoords(x, y, moveVector)
	for _, coords := range targetCoords {
		if matrix[coords.y][coords.x] == "[" || matrix[coords.y][coords.x] == "]" {
			// found another box
			canMove = append(canMove, canBoxMove(coords.x, coords.y, moveVector))
		} else if matrix[coords.y][coords.x] != "." {
			return false
		}
	}
	// if all true
	for _, move := range canMove {
		if !move {
			return false
		}
	}
	return true
}

func moveBox(x, y int, moveVector coord) {
	// boxes are 2 squares wide
	leftSide, rightSide := findBox(x, y)

	oldLeft := matrix[leftSide.y][leftSide.x]
	oldRight := matrix[rightSide.y][rightSide.x]

	// move both sides
	matrix[leftSide.y][leftSide.x] = "."
	matrix[rightSide.y][rightSide.x] = "."
	matrix[leftSide.y+moveVector.y][leftSide.x+moveVector.x] = oldLeft
	matrix[rightSide.y+moveVector.y][rightSide.x+moveVector.x] = oldRight

}

func loadData(filename string) {
	lines := framework.ReadInput(filename)

	parse := "matrix"
	for y, line := range lines {

		if line == "" {
			parse = "moves"
			continue
		}
		if parse == "matrix" {
			row := make([]string, 0)
			chars := strings.Split(line, "")
			for x, char := range chars {
				if char == "@" {
					robotX = x * 2
					robotY = y
					char = "."
				}
				if char == "O" {
					row = append(row, "[")
					row = append(row, "]")
				} else {
					row = append(row, char)
					row = append(row, char)
				}
			}
			matrix = append(matrix, row)
		}
		if parse == "moves" {
			// split moves into array
			moves = append(moves, strings.Split(line, "")...)
		}
	}

}
