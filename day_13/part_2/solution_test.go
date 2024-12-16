package main

import (
	"AoC2024/framework"
	"testing"
)

const correctAnswer = 74015623345775

func TestSolution(t *testing.T) {
	framework.RunTest(correctAnswer, solution(), t)
}