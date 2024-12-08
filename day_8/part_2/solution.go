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

				followVector(x1, y1, vector[0], vector[1])
				followVector(x1, y1, -vector[0], -vector[1])
				followVector(x2, y2, vector[0], vector[1])
				followVector(x2, y2, -vector[0], -vector[1])

			}
		}
	}

	total := countNodes()
	return total
}

func followVector(x, y int, vectorX, vectorY int) {
	var posX, posY int
	posX, posY = x, y

	for {
		// add vector
		posX += vectorX
		posY += vectorY

		// is outOfMatrix
		if posX < 0 || posX >= len(matrix[0]) || posY < 0 || posY >= len(matrix) {
			return
		}

		antiNodesMatrix[posY][posX] = 1
	}
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
