package main

import (
	"AoC2024/framework"
	"fmt"
	"strings"
)

var world [][]string
var visited [][]bool

type areaInfo struct {
	letter string
	area   int
	sides  int
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
				total += area.area * area.sides
			}
		}
	}
	return total
}

func grabArea(letter string, x, y int) areaInfo {
	var queue []point
	var regionPoints []point
	queue = append(queue, point{x: x, y: y})

	var area = 0
	var sides = 0
	for len(queue) > 0 {
		pointVal := queue[0]
		queue = queue[1:]
		if visited[pointVal.y][pointVal.x] == true {
			continue
		}

		visited[pointVal.y][pointVal.x] = true
		regionPoints = append(regionPoints, pointVal)
		area++
		sides += getCorners(letter, pointVal.x, pointVal.y)

		queue = append(queue, getNeighbors(letter, pointVal.x, pointVal.y)...)
	}

	return areaInfo{letter: letter, area: area, sides: sides}
}

func getCorners(letter string, x int, y int) int {
	corners := 0

	var n, e, w, s bool
	var ne, se, sw, nw bool

	if x > 0 && world[y][x-1] == letter {
		w = true
	}
	if x < len(world[0])-1 && world[y][x+1] == letter {
		e = true
	}
	if y > 0 && world[y-1][x] == letter {
		n = true
	}
	if y < len(world)-1 && world[y+1][x] == letter {
		s = true
	}

	if x > 0 && y > 0 && world[y-1][x-1] == letter {
		nw = true
	}
	if x < len(world[0])-1 && y > 0 && world[y-1][x+1] == letter {
		ne = true
	}
	if x < len(world[0])-1 && y < len(world)-1 && world[y+1][x+1] == letter {
		se = true
	}
	if x > 0 && y < len(world)-1 && world[y+1][x-1] == letter {
		sw = true
	}

	if !n && !w {
		corners++
	}
	if !n && !e {
		corners++
	}
	if !s && !e {
		corners++
	}
	if !s && !w {
		corners++
	}
	if n && w && !nw {
		corners++
	}
	if n && e && !ne {
		corners++
	}
	if s && e && !se {
		corners++
	}
	if s && w && !sw {
		corners++
	}

	return corners
}

// getNeighbors remains the same
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

// loadData remains the same
func loadData(filename string) {
	lines := framework.ReadInput(filename)
	world = nil
	visited = nil
	for _, line := range lines {
		world = append(world, strings.Split(line, ""))
		visited = append(visited, make([]bool, len(line)))
	}
}
