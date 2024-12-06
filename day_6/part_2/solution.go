package main

import (
	"AoC2024/framework"
	"fmt"
	"strconv"
	"strings"
)

var matrix [][]string

var guardX, guardY = 0, 0
var direction = "up"
var loops []string
var visitedCount [][]int

var initialDirection string
var initialGuardX, initialGuardY int

// how many times cell can be visited before it is considered a loop.
var loopDetectionThreshold = 4

func main() {
	total := solution()
	fmt.Println("solution:", total)
}

func solution() int {
	loadData()
	for y := 0; y < len(matrix); y++ {
		for x := 0; x < len(matrix[0]); x++ {
			initVisited()
			testRoute(x, y)
		}
	}
	return len(loops)
}

func testRoute(obstacleX, obstacleY int) {

	// if already obstacle, exit
	if matrix[obstacleY][obstacleX] == "#" {
		return
	}
	matrix[obstacleY][obstacleX] = "#"

	success := 1
	for success == 1 {
		success = moveGuard()

		if success == -1 {
			loops = append(loops, "obstacle: ("+strconv.Itoa(obstacleX)+","+strconv.Itoa(obstacleY)+")")
			break
		}
	}
	// clear obstacle after used
	matrix[obstacleY][obstacleX] = "."
}

func moveGuard() int {
	// check for obstacles

	canMove := 0
	for canMove == 0 {
		x, y := getDirection()
		newX := guardX + x
		newY := guardY + y
		canMove = canMoveTo(newX, newY)
		if canMove == -1 {
			return 0
		}
		if canMove == 1 {
			visitedCount[guardY][guardX]++

			guardX = newX
			guardY = newY
			if visitedCount[guardY][guardX] > loopDetectionThreshold {
				// loop
				return -1
			}

			return 1
		}
		if canMove == 0 {
			turn90()
		}
	}
	return 1
}

func canMoveTo(x, y int) int {
	if (x < 0) || (y < 0) || (x >= len(matrix)) || (y >= len(matrix[0])) {
		return -1
	}

	if matrix[y][x] == "#" {
		return 0
	}
	return 1
}

func getDirection() (x, y int) {
	if direction == "up" {
		return 0, -1
	} else if direction == "down" {
		return 0, 1
	} else if direction == "left" {
		return -1, 0
	} else {
		return 1, 0
	}
}

func turn90() {
	if direction == "up" {
		direction = "right"
	} else if direction == "right" {
		direction = "down"
	} else if direction == "down" {
		direction = "left"
	} else {
		direction = "up"
	}
}

func initVisited() {
	visitedCount = make([][]int, len(matrix))
	for i := range visitedCount {
		visitedCount[i] = make([]int, len(matrix[0]))
	}

	guardX = initialGuardX
	guardY = initialGuardY
	direction = initialDirection

}

func loadData() {
	lines := framework.ReadInput("input.txt")

	// init matrix
	matrix = make([][]string, 0)

	// split to matrix
	for _, line := range lines {
		matrix = append(matrix, strings.Split(line, ""))
	}

	// find guard initial position
	for y, line := range matrix {
		for x, char := range line {
			if char == "^" {
				initialGuardX = x
				initialGuardY = y
				initialDirection = "up"
				matrix[y][x] = "."
				break
			}
		}
	}
}
