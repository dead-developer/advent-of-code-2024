package main

import (
	"AoC2024/framework"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}
type machine struct {
	buttonA        point
	buttonB        point
	rewardLocation point
}

var machines []machine

func main() {
	total := solution()
	fmt.Println("solution:", total)
}

func solution() int {
	loadData("input.txt")

	total := 0

	for _, m := range machines {
		tokens := solveMachine(m)

		if tokens == -1 { // no solution
			continue
		}
		total += tokens
	}

	return total
}

func mustConvertToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

// GIPITI WROTE THIS SOLUTION
func solveMachine(m machine) int {
	Ax, Ay := m.buttonA.x, m.buttonA.y
	Bx, By := m.buttonB.x, m.buttonB.y
	Px, Py := m.rewardLocation.x, m.rewardLocation.y

	D := Ax*By - Ay*Bx

	if D != 0 {
		numA := Px*By - Py*Bx
		numB := Ax*Py - Ay*Px

		if numA%D != 0 || numB%D != 0 {
			return -1
		}

		a := numA / D
		b := numB / D

		if a < 0 || b < 0 {
			return -1
		}

		return 3*a + b

	} else {
		if Ax == 0 && Ay == 0 {

			if Px == 0 && Py == 0 {
				return 0
			}
			return -1
		}

		if Px*Ay != Py*Ax {
			return -1
		}

		// Check integral s:
		if Ax != 0 {
			if Px%Ax != 0 {
				return -1
			}
			s := Px / Ax

			if Ay != 0 && (Py%Ay != 0 || Py/Ay != s) {
				return -1
			}

			if Bx%Ax != 0 {
				return -1
			}
			k := Bx / Ax
			if Ay != 0 && (By%Ay != 0 || By/Ay != k) {
				return -1
			}
			return minimizeLinearEquation(s, k)

		} else {
			if Px != 0 {
				return -1
			}
			if Py%Ay != 0 {
				return -1
			}
			s := Py / Ay
			if Bx != 0 {
				return -1
			}
			if By%Ay != 0 {
				return -1
			}
			k := By / Ay
			return minimizeLinearEquation(s, k)
		}
	}
}

func minimizeLinearEquation(s, k int) int {
	if k == 0 {
		if s < 0 {
			return -1
		}
		return 3 * s
	}
	if k > 0 {
		if s < 0 {
			return -1
		}
		b := s / k
		a := s - k*b
		if a < 0 || b < 0 {
			return -1
		}
		return 3*a + b
	}
	ceilDiv := func(n, d int) int {
		return int(math.Ceil(float64(n) / float64(d)))
	}
	threshold := float64(s) / float64(k)
	var b int
	if threshold > 0 {
		b = ceilDiv(s, k)
	} else {
		b = 0
	}

	a := s - k*b
	if a < 0 || b < 0 {
		return -1
	}

	return 3*a + b
}

func loadData(filename string) {
	lines := framework.ReadInput(filename)

	reA := regexp.MustCompile(`Button A: X\+(\d+), Y\+(\d+)`)
	reB := regexp.MustCompile(`Button B: X\+(\d+), Y\+(\d+)`)
	price := regexp.MustCompile(`Prize: X=(\d+), Y=(\d+)`)

	var record machine
	// split to matrix
	for i, line := range lines {
		if line == "" {
			if i > 0 {
				machines = append(machines, record)
			}
			record = machine{}
		}

		if strings.HasPrefix(line, "Button A") {
			matches := reA.FindStringSubmatch(line)
			record.buttonA = point{mustConvertToInt(matches[1]), mustConvertToInt(matches[2])}
		}
		if strings.HasPrefix(line, "Button B") {
			matches := reB.FindStringSubmatch(line)
			record.buttonB = point{mustConvertToInt(matches[1]), mustConvertToInt(matches[2])}
		}
		if strings.HasPrefix(line, "Prize") {
			matches := price.FindStringSubmatch(line)
			record.rewardLocation = point{mustConvertToInt(matches[1]) + 10000000000000, mustConvertToInt(matches[2]) + 10000000000000}
		}

	}

	machines = append(machines, record)

}
