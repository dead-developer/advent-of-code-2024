package main

import (
	"AoC2024/framework"
	"testing"
)

const correctAnswer = 329

func TestSolution(t *testing.T) {
	framework.RunTest(correctAnswer, solution(), t)
}
