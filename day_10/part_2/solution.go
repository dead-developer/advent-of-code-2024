package main

import (
	"AoC2024/framework"
	"fmt"
	"strconv"
	"strings"
)

var topoMap [][]string

type trail struct {
	current string
	lookFor string
	x       int
	y       int
}
type point struct {
	x     int
	y     int
	value string
}

var queue []trail

func main() {
	total := solution()
	fmt.Println("solution:", total)
}

func solution() int {
	loadData()

	total := 0

	// find 0's and add to queue
	for y, line := range topoMap {
		for x, char := range line {
			if char == "0" {
				result := findTrail(x, y)
				total += result

				fmt.Println(result)
			}
		}
	}

	return total
}

func findTrail(x int, y int) int {
	result := 1

	queue = append(queue, trail{x: x, y: y, lookFor: "1", current: "0"})
	//find trail
	for len(queue) > 0 {
		//pop from queue
		trailVal := queue[0]
		queue = queue[1:]

		//mark as visited in trails

		// next letter
		val, _ := strconv.Atoi(trailVal.lookFor)
		findString := strconv.Itoa(val)
		//get neighbors
		neighbors := getNeighbors(trailVal.x, trailVal.y, findString)
		// increase val
		nextString := strconv.Itoa(val + 1)

		if len(neighbors) > 0 {
			result += len(neighbors) - 1
		} else if trailVal.current != "9" {
			result--
		}

		for _, neighbor := range neighbors {
			// has already been visited?
			queue = append(queue, trail{x: neighbor.x, y: neighbor.y, lookFor: nextString, current: findString})
		}
	}
	return result
}

func getNeighbors(x int, y int, findString string) []point {
	var neighbors []point

	if x > 0 && topoMap[y][x-1] == findString {
		neighbors = append(neighbors, point{x - 1, y, topoMap[y][x-1]})
	}

	if x < len(topoMap[0])-1 && topoMap[y][x+1] == findString {
		neighbors = append(neighbors, point{x + 1, y, topoMap[y][x+1]})
	}

	if y > 0 && topoMap[y-1][x] == findString {
		neighbors = append(neighbors, point{x, y - 1, topoMap[y-1][x]})
	}
	if y < len(topoMap)-1 && topoMap[y+1][x] == findString {
		neighbors = append(neighbors, point{x, y + 1, topoMap[y+1][x]})
	}

	return neighbors
}

func loadData() {
	lines := framework.ReadInput("input.txt")

	for _, line := range lines {
		topoMap = append(topoMap, strings.Split(line, ""))
	}

}
