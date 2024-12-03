package main

import (
	"AoC2024/framework"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var line string

func main() {
	total := solution()
	fmt.Println("solution:", total)
}

func solution() int {
	loadData()

	total := 0

	re, _ := regexp.Compile(`mul\((\d+),(\d+)\)|do\(\)|don't\(\)`)

	enabled := true
	for _, match := range re.FindAllStringSubmatch(line, -1) {
		if match[0] == "do()" {
			enabled = true
			continue
		}
		if match[0] == "don't()" {
			enabled = false
			continue
		}

		if enabled {
			if len(match) == 3 {
				num1, _ := strconv.Atoi(match[1])
				num2, _ := strconv.Atoi(match[2])

				total += num1 * num2
			}
		}

		fmt.Println(match[0])
	}

	return total
}

func loadData() {
	lines := framework.ReadInput("input.txt")

	line = strings.Join(lines, "")

}
