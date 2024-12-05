package main

import (
	"AoC2024/framework"
	"fmt"
	"strconv"
	"strings"
)

var rules [][]int
var prints [][]int

func main() {
	total := solution()
	fmt.Println("solution:", total)
}

func solution() int {
	loadData()

	total := 0

	for _, pages := range prints {
		total += checkPrints(pages)
	}

	return total
}

func checkPrints(pages []int) int {
	if !checkPage(pages) {
		return 0
	}
	//return middle number
	return pages[len(pages)/2]

}

func checkPage(pages []int) bool {
	// check if pages follow rules
	for _, rule := range rules {
		beforeIndex := findInSlice(pages, rule[0])
		afterIndex := findInSlice(pages, rule[1])
		// either page not included
		if beforeIndex == -1 || afterIndex == -1 {
			continue
		}
		// is page in rule[0] is before rule[1]
		if beforeIndex > afterIndex {
			return false
		}
	}

	return true
}

func findInSlice(slice []int, target int) int {
	for i, v := range slice {
		if v == target {
			return i
		}
	}
	return -1
}

func loadData() {
	lines := framework.ReadInput("input.txt")

	section := 0

	// parse rules until empty line
	for _, line := range lines {
		if line == "" {
			section = 1
			continue
		}
		if section == 0 {
			parts := strings.Split(line, "|")
			before, _ := strconv.Atoi(parts[0])
			after, _ := strconv.Atoi(parts[1])
			rules = append(rules, []int{before, after})
		}
		if section == 1 {
			// parse prints
			parts := strings.Split(line, ",")
			// convert parts to int
			pages := make([]int, 0)
			for _, part := range parts {
				page, _ := strconv.Atoi(part)
				pages = append(pages, page)
			}
			prints = append(prints, pages)
		}
	}

}
