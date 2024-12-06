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

func main() {
	total := solution()
	fmt.Println("solution:", total)
}

func solution() int {
	loadData()
	initVisitedCounts()
	for y := 0; y < len(matrix); y++ {
		for x := 0; x < len(matrix[0]); x++ {
			loadData()

			testRoute(x, y)
		}
	}
	return len(loops)
}

func testRoute(obstacleX, obstacleY int) {

	matrix[obstacleY][obstacleX] = "#"
	loopThreshold := len(matrix) * len(matrix[0])
	success := true
	for success {
		success = moveGuard()
		loopThreshold--
		if loopThreshold == 0 {
			loops = append(loops, "obstacle: ("+strconv.Itoa(obstacleX)+","+strconv.Itoa(obstacleY)+")")
			break
		}
	}
}

func moveGuard() bool {
	// check for obstacles

	canMove := 0
	for canMove == 0 {
		x, y := getDirection()
		newX := guardX + x
		newY := guardY + y
		canMove = canMoveTo(newX, newY)
		if canMove == -1 {
			return false
		}
		if canMove == 1 {
			guardX = newX
			guardY = newY
			//matrix[guardY][guardX] = "X"
			visitedCount[guardY][guardX]++
			return true
		}
		if canMove == 0 {
			turn90()
		}
	}
	return true
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

func visualize() {
	for _, line := range matrix {
		fmt.Println(line)
	}
}

func initVisitedCounts() {
	visitedCount = make([][]int, len(matrix))
	for i := range visitedCount {
		visitedCount[i] = make([]int, len(matrix[0]))
	}
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
				guardX = x
				guardY = y
				direction = "up"
				break
			}
		}
	}

	matrix[guardY][guardX] = "X"
}
