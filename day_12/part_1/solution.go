package main

import (
	"AoC2024/framework"
	"fmt"
	"strings"
)

var world [][]string
var visited [][]bool

type areaInfo struct {
	letter    string
	area      int
	perimeter int
}

type point struct {
	x int
	y int
}

func main() {
	total := solution()
	fmt.Println("solution:", total)
}

func solution() int {
	loadData("input.txt")
	total := 0
	for y := 0; y < len(world); y++ {
		for x := 0; x < len(world[0]); x++ {
			if visited[y][x] == false {
				area := grabArea(world[y][x], x, y)

				fmt.Println(area)

				total += area.area * area.perimeter
			}
		}
	}

	return total
}

func grabArea(letter string, x, y int) areaInfo {
	var queue []point
	queue = append(queue, point{x: x, y: y})

	var area = 0
	var perimeter int
	for len(queue) > 0 {
		//pop from queue
		pointVal := queue[0]
		queue = queue[1:]
		if visited[pointVal.y][pointVal.x] == true {
			continue
		}

		visited[pointVal.y][pointVal.x] = true

		area++
		perimeter += getPerimeters(letter, pointVal.x, pointVal.y)
		queue = append(queue, getNeighbors(letter, pointVal.x, pointVal.y)...)

	}
	return areaInfo{letter: letter, area: area, perimeter: perimeter}
}

func getPerimeters(letter string, x int, y int) int {
	var perimeter = 4
	if x > 0 && world[y][x-1] == letter {
		perimeter--
	}
	if x < len(world[0])-1 && world[y][x+1] == letter {
		perimeter--
	}
	if y > 0 && world[y-1][x] == letter {
		perimeter--
	}
	if y < len(world)-1 && world[y+1][x] == letter {
		perimeter--
	}
	return perimeter
}

func getNeighbors(letter string, x int, y int) []point {
	var neighbors []point
	if x > 0 && world[y][x-1] == letter && !visited[y][x-1] {
		neighbors = append(neighbors, point{x - 1, y})
	}
	if x < len(world[0])-1 && world[y][x+1] == letter && !visited[y][x+1] {
		neighbors = append(neighbors, point{x + 1, y})
	}
	if y > 0 && world[y-1][x] == letter && !visited[y-1][x] {
		neighbors = append(neighbors, point{x, y - 1})
	}
	if y < len(world)-1 && world[y+1][x] == letter && !visited[y+1][x] {
		neighbors = append(neighbors, point{x, y + 1})
	}
	return neighbors
}

func loadData(filename string) {
	lines := framework.ReadInput(filename)

	// split to matrix
	for _, line := range lines {
		world = append(world, strings.Split(line, ""))
		visited = append(visited, make([]bool, len(line)))
	}

}
