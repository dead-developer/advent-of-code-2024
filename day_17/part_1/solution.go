package main

import (
	"AoC2024/framework"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

var program []int
var registers = make(map[string]int)

var commandPointer int

var output []int

func main() {
	total := solution()

	fmt.Println("solution:", total)
}

func solution() string {
	loadData("input.txt")

	for commandPointer < len(program) {
		command, _ := getOpCode()
		operand, _ := getOperand(command)

		processCode(command, operand)
	}

	outputStrings := make([]string, len(output))
	for i, num := range output {
		outputStrings[i] = strconv.Itoa(num)
	}
	outputString := strings.Join(outputStrings, ",")

	return outputString
}

func getOpCode() (int, bool) {
	if commandPointer < len(program) {
		commandPointer++
		return program[commandPointer-1], true
	}
	return 0, false
}

func getOperand(command int) (int, bool) {
	operandType := "normal"
	if command == 0 || command == 2 || command == 5 || command == 6 || command == 7 {
		operandType = "combo"
	}
	operand := 0
	if operandType == "normal" {
		return getOpCode()
	}

	operand, _ = getOpCode()
	if operand == 4 {
		return registers["A"], true
	}
	if operand == 5 {
		return registers["B"], true
	}
	if operand == 6 {
		return registers["C"], true
	}
	if operand == 7 {
		panic("invalid program")
	}

	return operand, true
}

func processCode(command int, operand int) {
	if command == 0 {
		// adv
		denominator := math.Pow(2, float64(operand))
		i := registers["A"] / int(denominator)
		registers["A"] = i
	}
	if command == 1 {
		// bxl bitwise XOR
		registers["B"] = registers["B"] ^ operand
	}
	if command == 2 {
		// bst - module 8
		registers["B"] = operand % 8
	}
	if command == 3 {
		// jnz - jump if not zero
		if registers["A"] > 0 {
			commandPointer = operand
		}
	}
	if command == 4 {
		// bxc - bitwise XOR
		registers["B"] = registers["B"] ^ registers["C"]
	}
	if command == 5 {
		// out
		output = append(output, operand%8)
	}
	if command == 6 {
		// bdv
		denominator := math.Pow(2, float64(operand))
		i := registers["A"] / int(denominator)
		registers["B"] = i
	}
	if command == 7 {
		// bdv
		denominator := math.Pow(2, float64(operand))
		i := registers["A"] / int(denominator)
		registers["C"] = i
	}
}

func loadData(filename string) {
	lines := framework.ReadInput(filename)

	reA, _ := regexp.Compile(`Register (.+): (\d+)`)
	for _, line := range lines {
		if strings.HasPrefix(line, "Register") {
			matches := reA.FindStringSubmatch(line)
			registers[matches[1]] = mustConvertToInt(matches[2])
			continue
		}

		if strings.HasPrefix(line, "Program") {
			parts := strings.Split(line, ":")
			opCodes := strings.Split(parts[1], ",")
			program = make([]int, len(opCodes))
			// convert to integers
			for i, code := range opCodes {
				program[i] = mustConvertToInt(code)
			}
			continue
		}
	}
}

func mustConvertToInt(s string) int {
	s = strings.TrimSpace(s)
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
