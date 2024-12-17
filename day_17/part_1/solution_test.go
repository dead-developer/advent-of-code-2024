package main

import (
	"AoC2024/framework"
	"testing"
)

const correctAnswer = "7,3,1,3,6,3,6,0,2"

func TestSolution(t *testing.T) {
	framework.RunTestString(correctAnswer, solution(), t)
}
