package main

import (
	"AoC2024/framework"
	"fmt"
	"regexp"
	"strconv"
)

const areaSizeX = 101
const areaSizeY = 103

var matrix [][]int
var quadrants = make(map[int][]robot)

type robot struct {
	x, y                 int
	velocityX, velocityY int
}

var robots []robot

func main() {
	total := solution()
	fmt.Println("solution:", total)
}

func solution() int {
	loadData("input.txt")

	for j := 0; j < areaSizeY*areaSizeX; j++ {
		generateForTurn(1) // step 1 by 1

		if findLines(10) { // find long vertical line
			//visualize(j + 1)
			return j + 1
		}
	}

	return 0
}

func generateForTurn(turn int) {
	resetImage()
	for i, robot := range robots {
		robots[i].x, robots[i].y = moveRobot(robot, turn)
		matrix[robots[i].y][robots[i].x]++
	}
}

func findLines(threshold int) bool {
	for y := 0; y < areaSizeY-threshold+1; y++ {
		for x := 0; x < areaSizeX; x++ {
			if matrix[y][x] > 0 {
				found := true
				for i := 0; i < threshold; i++ {
					if matrix[y+i][x] == 0 {
						found = false
					}
				}
				if found {
					return true
				}
			}
		}
	}
	return false
}

func visualize(turn int) {
	fmt.Println()
	fmt.Println("Turn", turn)
	for _, row := range matrix {
		for _, num := range row {
			if num == 0 {
				fmt.Print(".")
			} else {
				fmt.Print(strconv.Itoa(num))
			}
		}
		fmt.Println()
	}
}

func moveRobot(robot robot, turn int) (int, int) {
	var newX, newY int
	newX = applyTeleport(areaSizeX, robot.x+(robot.velocityX)*turn)
	newY = applyTeleport(areaSizeY, robot.y+(robot.velocityY)*turn)
	return newX, newY
}

func applyTeleport(size, value int) int {
	return (value%size + size) % size
}

func mustConvertToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func resetImage() {
	matrix = make([][]int, areaSizeY)
	for i := range matrix {
		matrix[i] = make([]int, areaSizeX)
	}
}

func loadData(filename string) {
	lines := framework.ReadInput(filename)
	re := regexp.MustCompile(`p=(-?\d+),(-?\d+) v=(-?\d+),(-?\d+)`)
	for _, line := range lines {
		matches := re.FindStringSubmatch(line)

		robots = append(robots, robot{
			mustConvertToInt(matches[1]),
			mustConvertToInt(matches[2]),
			mustConvertToInt(matches[3]),
			mustConvertToInt(matches[4]),
		})
	}
}
