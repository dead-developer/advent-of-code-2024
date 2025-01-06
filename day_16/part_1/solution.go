package main

import (
	"AoC2024/framework"
	"fmt"
	"math"
	"strings"
)

type cell struct {
	cost      int
	direction point
}

var maze [][]string
var visited = make(map[point]bool)
var nodes = make(map[point]cell)
var parents = make(map[point]point)

var nilPoint = point{-1, -1}

var startPoint, endPoint point

type point struct {
	x, y int
}

func main() {
	total := solution()

	fmt.Println("solution:", total)
}

func solution() int {
	loadData("input.txt")

	nodes[startPoint] = cell{cost: 0, direction: point{1, 0}}

	nextNodeId := getNextNode()
	for nextNodeId != nilPoint {
		processNode(nextNodeId)

		nextNodeId = getNextNode()
		if nextNodeId == endPoint {
			break
		}
	}

	return nodes[endPoint].cost
}

func getNeighbours(node point) []point {
	var neighbors []point
	if node.x-1 >= 0 && maze[node.y][node.x-1] != "#" {
		neighbors = append(neighbors, point{node.x - 1, node.y})
	}
	if node.x+1 < len(maze[0]) && maze[node.y][node.x+1] != "#" {
		neighbors = append(neighbors, point{node.x + 1, node.y})
	}
	if node.y-1 >= 0 && maze[node.y-1][node.x] != "#" {
		neighbors = append(neighbors, point{node.x, node.y - 1})
	}
	if node.y+1 < len(maze) && maze[node.y+1][node.x] != "#" {
		neighbors = append(neighbors, point{node.x, node.y + 1})
	}
	return neighbors
}

func processNode(nodeId point) {
	visited[nodeId] = true
	for _, neighbour := range getNeighbours(nodeId) {
		addNode(neighbour, nodeId)
	}
}

func addNode(nodeId point, parentNode point) {
	// if already exists
	currentDirection := nodes[parentNode].direction

	if parentNode == nilPoint {
		currentDirection = point{1, 0}
	}
	newDirection := point{nodeId.x - parentNode.x, nodeId.y - parentNode.y}
	newCost := 1
	if currentDirection != newDirection {
		newCost = 1001
	}

	newCost += nodes[parentNode].cost

	if _, ok := nodes[nodeId]; ok {
		// if old cost + new cost < current cost
		if newCost < nodes[nodeId].cost {
			nodes[nodeId] = cell{
				cost:      newCost,
				direction: newDirection,
			}
			parents[parentNode] = nodeId
		}
	} else {
		nodes[nodeId] = cell{
			cost:      newCost,
			direction: newDirection,
		}
		parents[parentNode] = nodeId
	}

}

func getNextNode() point {
	var lowestCost = math.MaxInt
	var lowestCostNode = nilPoint
	for nodeId, node := range nodes {
		if visited[nodeId] {
			continue
		}
		if node.cost < lowestCost {
			lowestCost = node.cost
			lowestCostNode = nodeId
		}
	}
	return lowestCostNode
}

func loadData(filename string) {
	lines := framework.ReadInput(filename)

	maze = make([][]string, 0)
	for _, line := range lines {
		maze = append(maze, strings.Split(line, ""))
	}

	// find start and end locations
	for y, line := range maze {
		for x, char := range line {
			if char == "S" {
				startPoint = point{x: x, y: y}
				maze[y][x] = "."
				continue
			}
			if char == "E" {
				endPoint = point{x: x, y: y}
				maze[y][x] = "."
				continue
			}
		}
	}

}
