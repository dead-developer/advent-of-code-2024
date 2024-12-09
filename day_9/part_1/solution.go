package main

import (
	"AoC2024/framework"
	"fmt"
	"strconv"
	"strings"
)

var data []string

var blocks []string

func main() {
	total := solution()
	fmt.Println("solution:", total)
}

func solution() int {
	loadData()

	unpack()

	lastIndex := len(blocks) - 1
	for i := 0; i < len(blocks); i++ {
		if blocks[i] == " " {
			break
		}
		if blocks[i] == "." {
			//move last to current
			blocks[i] = blocks[lastIndex]
			blocks[lastIndex] = " "
			lastIndex--
			// remove trailing empties
			for j := lastIndex; j >= 0; j-- {
				if blocks[j] != "." {
					break
				}
				blocks[j] = " "
				lastIndex--
			}
		}

	}

	total := 0
	for i := 0; i < lastIndex+1; i++ {
		value, _ := strconv.Atoi(blocks[i])
		total += i * value
	}

	return total
}

func unpack() {
	currentType := "file"
	fileId := 0
	for i := 0; i < len(data); i++ {
		amount, _ := strconv.Atoi(data[i])
		if currentType == "file" {
			for j := 0; j < amount; j++ {
				idName := strconv.Itoa(fileId)
				blocks = append(blocks, idName)
			}
			currentType = "folder"
			fileId++
		} else {
			for j := 0; j < amount; j++ {
				blocks = append(blocks, ".")
			}
			currentType = "file"
		}

	}
}

func loadData() {
	lines := framework.ReadInput("input.txt")

	data = strings.Split(lines[0], "")

}
