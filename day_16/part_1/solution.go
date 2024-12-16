package main

import (
	"AoC2024/framework"
	"fmt"
	"sort"
	"strings"
)

var maze [][]string

var found []int

type point struct {
	x, y int
}

type reindeerStruct struct {
	position  point
	cost      int
	direction int
	visited   map[point]bool
}

var start = point{}
var theEnd = point{}

var directions = map[int]point{
	0: {0, -1}, //N
	1: {1, 0},  //E
	2: {0, 1},  //S
	3: {-1, 0}, //W
}

var queue []queueItem

type queueItem struct {
	reindeer      reindeerStruct
	goToDirection int
}

func main() {
	total := solution()

	fmt.Println("solution:", total)
}

func solution() int {
	loadData("input.txt")

	total := 0

	firstReindeer := reindeerStruct{start, 0, 1, make(map[point]bool)}
	firstReindeer.visited[start] = true
	queue = append(queue, queueItem{reindeer: firstReindeer, goToDirection: 1}) // any

	for len(queue) > 0 {
		reindeerItem := queue[0]
		queue = queue[1:]

		cost := processPath(reindeerItem.reindeer.position, &reindeerItem.reindeer)
		if cost > 0 {
			found = append(found, cost)
		}

	}
	// sort found
	sort.Ints(found)
	fmt.Println(found)
	return total
}

func processPath(startPos point, reindeer *reindeerStruct) int {
	// where can I move?

	for true {

		availableDirections := getDirections(reindeer.position, reindeer)

		isEnd := theEnd.x == reindeer.position.x && theEnd.y == reindeer.position.y
		if isEnd {
			return reindeer.cost
		}
		if len(availableDirections) == 0 { // dead end
			return -1
		}
		firstDirection := availableDirections[0]
		availableDirections = availableDirections[1:]
		MoveTo(reindeer, firstDirection)

		//spawn more reindeers for other paths
		for _, followDirection := range availableDirections {
			// copy reindeer
			newReindeer := reindeerStruct{reindeer.position, reindeer.cost, reindeer.direction, reindeer.visited}
			queue = append(queue, queueItem{reindeer: newReindeer, goToDirection: followDirection})
		}
	}
	return 0
}

func MoveTo(reindeer *reindeerStruct, newDirection int) {

	// if not facing the right way, turn
	turnReindeer(newDirection, reindeer)

	// move
	reindeer.position.x += directions[newDirection].x
	reindeer.position.y += directions[newDirection].y
	reindeer.visited[reindeer.position] = true
	reindeer.cost += 1

}

func turnReindeer(towards int, reindeer *reindeerStruct) {
	if towards == reindeer.direction {
		return
	}
	diff := (towards - reindeer.direction + 4) % 4
	switch diff {
	case 1:
		reindeer.direction = (reindeer.direction + 1) % 4
		reindeer.cost += 1000
		break
	case 2:
		reindeer.direction = (reindeer.direction + 3) % 4
		reindeer.direction = (reindeer.direction + 3) % 4
		reindeer.cost += 2000
		break
	case 3:
		reindeer.direction = (reindeer.direction + 3) % 4
		reindeer.cost += 1000
		break
	}

}

func getDirections(location point, reindeer *reindeerStruct) []int {
	var availableDirections []int

	for direction, directionVector := range directions {
		if maze[location.y+directionVector.y][location.x+directionVector.x] != "#" {
			if isVisited(point{location.x + directionVector.x, location.y + directionVector.y}, reindeer) {
				continue
			}
			availableDirections = append(availableDirections, direction)
		}
	}
	return availableDirections
}

func isVisited(location point, reindeer *reindeerStruct) bool {
	if _, ok := reindeer.visited[location]; ok {
		return true
	}
	return false
}

func loadData(filename string) {
	lines := framework.ReadInput(filename)

	maze = make([][]string, 0)
	for _, line := range lines {
		maze = append(maze, strings.Split(line, ""))
	}

	// find start S and end
	for y, line := range maze {
		for x, char := range line {
			if char == "S" {
				start = point{x: x, y: y}
				maze[y][x] = "."
				continue
			}
			if char == "E" {
				theEnd = point{x: x, y: y}
				maze[y][x] = "."
				continue
			}
		}
	}

}
