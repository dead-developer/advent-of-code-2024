package main

import (
	"AoC2024/framework"
	"fmt"

	"regexp"
	"strconv"
)

const areaSizeX = 101
const areaSizeY = 103

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

	total := 0

	for i, robot := range robots {
		robots[i].x, robots[i].y = moveRobot(robot, 100)

	}

	splitToQuadrants()

	for _, robots := range quadrants {
		if total == 0 {
			total = len(robots)
			continue
		}
		total = total * len(robots)
	}

	return total
}

func splitToQuadrants() {
	halfX := areaSizeX / 2
	halfY := areaSizeY / 2

	for _, robot := range robots {
		// place in area

		if robot.x < halfX && robot.y < halfY {
			quadrants[1] = append(quadrants[1], robot)
		} else if robot.x > halfX && robot.y < halfY {
			quadrants[2] = append(quadrants[2], robot)

		} else if robot.x < halfX && robot.y > halfY {
			quadrants[3] = append(quadrants[3], robot)

		} else if robot.x > halfX && robot.y > halfY {
			quadrants[4] = append(quadrants[4], robot)
		}
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
