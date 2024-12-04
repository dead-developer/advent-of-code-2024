package main

import (
	"AoC2024/framework"
	"testing"
)

const correctAnswer = 1877

func TestSolution(t *testing.T) {
	framework.RunTest(correctAnswer, solution(), t)
}
