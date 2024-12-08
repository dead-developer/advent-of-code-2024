package main

import (
	"AoC2024/framework"
	"fmt"
	"strings"
)

type antenna struct {
	x, y int
	freq string
}
type antiNode struct {
	x, y int
}

var matrix [][]string
var antiNodesMatrix [][]int

var antennas []antenna
var frequencies = make(map[string]int)

func main() {
	total := solution()
	fmt.Println("solution:", total)
}

func solution() int {
	loadData()

	for freq, count := range frequencies {
		if count > 1 {
			antennasOfFreq := antennasOfFreq(freq)

			pairs := pairs(antennasOfFreq)

			for _, pair := range pairs {
				x1, y1 := pair[0].x, pair[0].y
				x2, y2 := pair[1].x, pair[1].y
				vector := distanceVector(x1, y1, x2, y2)

				// add to first node
				var x, y int

				x, y = calcNode(pair[0], vector, false)
				if isLegalPosition(pairs, x, y) {
					antiNodesMatrix[y][x] = 1
				}

				x, y = calcNode(pair[0], vector, true)
				if isLegalPosition(pairs, x, y) {
					antiNodesMatrix[y][x] = 1
				}

				x, y = calcNode(pair[1], vector, false)
				if isLegalPosition(pairs, x, y) {
					antiNodesMatrix[y][x] = 1
				}

				x, y = calcNode(pair[1], vector, true)
				if isLegalPosition(pairs, x, y) {
					antiNodesMatrix[y][x] = 1
				}
			}
		}
	}
	total := countNodes()
	return total
}

func calcNode(node antenna, distVector []int, negative bool) (int, int) {
	if negative {
		//negative vector
		return node.x - distVector[0], node.y - distVector[1]
	}
	return node.x + distVector[0], node.y + distVector[1]
}

func countNodes() int {
	count := 0
	for _, row := range antiNodesMatrix {
		for _, col := range row {
			if col == 1 {
				count++
			}
		}
	}
	return count
}

func isLegalPosition(pairs [][]antenna, x, y int) bool {
	// ignore if either of pairs
	for _, pair := range pairs {
		if pair[0].x == x && pair[0].y == y {
			return false
		}
		if pair[1].x == x && pair[1].y == y {
			return false
		}
	}

	return x >= 0 && x < len(matrix[0]) && y >= 0 && y < len(matrix)
}

func antennasOfFreq(freq string) []antenna {
	antennasOfFreq := make([]antenna, 0)
	for _, antenna := range antennas {
		if antenna.freq == freq {
			antennasOfFreq = append(antennasOfFreq, antenna)
		}
	}
	return antennasOfFreq
}

func distanceVector(x1, y1, x2, y2 int) []int {
	return []int{x1 - x2, y1 - y2}
}

func pairs(nodes []antenna) [][]antenna {
	// generate pair pairs from nodes
	var pairs [][]antenna
	for i := 0; i < len(nodes); i++ {
		for j := i + 1; j < len(nodes); j++ {
			pairs = append(pairs, []antenna{nodes[i], nodes[j]})
		}
	}
	return pairs
}

func visualize() {
	for _, row := range matrix {
		fmt.Println(row)
	}
	fmt.Println()

	for _, row := range antiNodesMatrix {
		fmt.Println(row)
	}
	fmt.Println()

	fmt.Println(antennas)
}

func loadData() {
	lines := framework.ReadInput("input.txt")

	// split to matrix
	for _, line := range lines {
		matrix = append(matrix, strings.Split(line, ""))
	}

	// init antiNodesMatrix with 0
	antiNodesMatrix = make([][]int, 0)
	for _, line := range lines {
		antiNodesMatrix = append(antiNodesMatrix, make([]int, len(line)))
	}

	// find antennas
	for y, line := range matrix {
		for x, char := range line {
			if char != "." {
				antennas = append(antennas, antenna{x, y, char})
				frequencies[char]++
			}
		}
	}
}
