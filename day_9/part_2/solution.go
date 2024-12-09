package main

import (
	"AoC2024/framework"
	"fmt"
	"strconv"
	"strings"
)

var data []string

var blocks []string

type file struct {
	id    string
	index int
	size  int
}

var files []file

func main() {
	total := solution()
	fmt.Println("solution:", total)
}

func solution() int {
	loadData()

	unpack()

	for i := len(files) - 1; i > 0; i-- {
		fitIndex := findFit(files[i].index, files[i].size)
		if fitIndex > 0 {
			moveFile(files[i].id, fitIndex, files[i].index, files[i].size)
		}

	}
	total := 0
	for i := 0; i < len(blocks); i++ {
		if blocks[i] == "." {
			continue
		}
		value, _ := strconv.Atoi(blocks[i])
		total += i * value
	}
	return total
}

func findFit(maxIndex int, size int) int {
	for i := 0; i < maxIndex; i++ {
		if blocks[i] == "." {
			space := calcEmptySpace(i)
			if space >= size {
				return i
			}
		}
	}
	return 0
}

func moveFile(id string, targetIndex int, sourceIndex int, size int) {
	for i := 0; i < size; i++ {
		blocks[targetIndex+i] = id
		blocks[sourceIndex+i] = "."
	}
}

func calcEmptySpace(index int) int {
	size := 0

	for i := index; i < len(blocks); i++ {
		if blocks[i] != "." {
			break
		}
		size++
	}
	return size
}

func unpack() {
	currentType := "file"
	fileId := 0
	index := 0
	for i := 0; i < len(data); i++ {
		amount, _ := strconv.Atoi(data[i])
		if currentType == "file" {
			idName := strconv.Itoa(fileId)
			for j := 0; j < amount; j++ {
				blocks = append(blocks, idName)
			}
			files = append(files, file{id: idName, size: amount, index: index})
			currentType = "folder"
			fileId++
		} else {
			for j := 0; j < amount; j++ {
				blocks = append(blocks, ".")
			}
			currentType = "file"
		}

		index += amount
	}
}

func loadData() {
	lines := framework.ReadInput("input.txt")

	data = strings.Split(lines[0], "")

}
