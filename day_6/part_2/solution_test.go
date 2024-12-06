package main

import (
	"AoC2024/framework"
	"testing"
)

const correctAnswer = 1503

func TestSolution(t *testing.T) {
	framework.RunTest(correctAnswer, solution(), t)
}
