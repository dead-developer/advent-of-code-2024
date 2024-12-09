package main

import (
	"AoC2024/framework"
	"testing"
)

const correctAnswer = 6250605700557

func TestSolution(t *testing.T) {
	framework.RunTest(correctAnswer, solution(), t)
}
