//go:build part1

package main

import (
	"os"
	"testing"
)

var input1 string
var anchor1 int

func init() {
	data, _ := os.ReadFile("../input/day3.txt")
	input1 = string(data)
}

func BenchmarkPart1(b *testing.B) {
	var result int
	for b.Loop() {
		Answer = 0
		solve(input1)
		result = Answer
	}
	anchor1 = result
}
