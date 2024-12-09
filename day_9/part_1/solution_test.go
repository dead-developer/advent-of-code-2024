package main

import (
	"AoC2024/framework"
	"testing"
)

const correctAnswer = 6225730762521

func TestSolution(t *testing.T) {
	framework.RunTest(correctAnswer, solution(), t)
}
