package main

import (
	"AoC2024/framework"
	"fmt"
	"strings"
)

var maze [][]string
var visited map[point]bool

var startPoint, endPoint point

var queue []queueItem

type queueItem struct {
	cell     point
	sourceId point
	cost     int
}

type point struct {
	x, y int
}

type edge struct {
	from point
	to   point
	cost int
}

type node struct {
	id       point
	edges    map[point]*edge
	sourceId point
	label    string
}

var nodes = make(map[point]*node)

var nilPoint = point{-1, -1}

var directions = map[int]point{
	0: {0, -1}, //N
	1: {1, 0},  //E
	2: {0, 1},  //S
	3: {-1, 0}, //W
}

func main() {
	total := solution()

	fmt.Println("solution:", total)
}

func solution() int {
	loadData("input.txt")
	buildGraph()

	return 0
}

//
//func turnReindeer(towards int, reindeer *reindeerStruct) {
//	if towards == reindeer.direction {
//		return
//	}
//	diff := (towards - reindeer.direction + 4) % 4
//	switch diff {
//	case 1:
//		reindeer.direction = (reindeer.direction + 1) % 4
//		reindeer.cost += 1000
//		break
//	case 2:
//		reindeer.direction = (reindeer.direction + 3) % 4
//		reindeer.direction = (reindeer.direction + 3) % 4
//		reindeer.cost += 2000
//		break
//	case 3:
//		reindeer.direction = (reindeer.direction + 3) % 4
//		reindeer.cost += 1000
//		break
//	}
//
//}

func getNeighbours(location point) []point {
	var availableDirections []point

	for _, directionVector := range directions {
		location := point{location.x + directionVector.x, location.y + directionVector.y}
		if maze[location.y][location.x] != "#" {
			if visited[location] {
				continue
			}
			availableDirections = append(availableDirections, location)
		}
	}
	return availableDirections
}

func buildGraph() {
	// create starting node
	visited = make(map[point]bool)
	queue = make([]queueItem, 0)

	//nodes[startPoint] = node{id: startPoint, edges: make(map[int]int), sourceId: nilPoint}
	queue = append(queue, queueItem{cell: startPoint, sourceId: nilPoint})

	for len(queue) > 0 {
		currentQueueItem := queue[0]
		queue = queue[1:]

		visited[currentQueueItem.cell] = true
		// add node
		newNode := node{id: currentQueueItem.cell, edges: make(map[point]*edge), sourceId: currentQueueItem.sourceId}
		if currentQueueItem.cell == endPoint {
			newNode.label = "E"
		}
		nodes[currentQueueItem.cell] = &newNode

		neighbours := getNeighbours(currentQueueItem.cell)
		for _, neighbour := range neighbours {
			// add edge to source node
			if currentQueueItem.sourceId != nilPoint {
				newEdge := edge{from: currentQueueItem.sourceId, to: neighbour, cost: currentQueueItem.cost}
				nodes[currentQueueItem.cell].edges[neighbour] = &newEdge
			}
			queue = append(queue, queueItem{cell: neighbour, sourceId: currentQueueItem.cell, cost: 1})

		}
		// neighbours are empty, remove parent node
	}
	optimizeGraph()
	visualize()

}

func optimizeGraph() {
	// remove deadEnds
	for nodeId := range nodes {
		checkDeadEnd(nodeId)
	}
}

func checkDeadEnd(nodeId point) {
	if nodeId == nilPoint {
		return
	}
	if _, ok := nodes[nodeId]; ok {
		currentNode := nodes[nodeId]
		if currentNode.label == "E" {
			return
		}
		if len(currentNode.edges) == 0 {
			deleteNode(nodeId)

			// if dead end found, check source node
			checkDeadEnd(currentNode.sourceId)
			return
		}
	}
}

func deleteNode(nodeId point) {
	if _, ok := nodes[nodeId]; ok {
		// remove node and source nodes edge
		currentNode := nodes[nodeId]
		if currentNode.sourceId != nilPoint {
			delete(nodes[currentNode.sourceId].edges, nodeId)
		}
		delete(nodes, nodeId)
	}
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

func visualize() {
	for y := 0; y < len(maze); y++ {
		for x := 0; x < len(maze[y]); x++ {
			fmt.Print(maze[y][x])
		}
		fmt.Println()
	}
	fmt.Println()

	// print individiual nores
	for _, node := range nodes {
		fmt.Println(node.id, node.edges)
	}

}

func waitForKeyPress() {
	_, _ = fmt.Scanln()
}
