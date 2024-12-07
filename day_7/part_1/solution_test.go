package main

import (
	"AoC2024/framework"
	"testing"
)

const correctAnswer = 1399219271639

func TestSolution(t *testing.T) {
	framework.RunTest(correctAnswer, solution(), t)
}
