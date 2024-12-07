package main

import (
	"AoC2024/framework"
	"testing"
)

const correctAnswer = 275791737999003

func TestSolution(t *testing.T) {
	framework.RunTest(correctAnswer, solution(), t)
}
