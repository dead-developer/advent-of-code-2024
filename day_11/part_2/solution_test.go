package main

import (
	"AoC2024/framework"
	"testing"
)

const correctAnswer = 257246536026785

func TestSolution(t *testing.T) {
	framework.RunTest(correctAnswer, solution(), t)
}
