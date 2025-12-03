//go:build part2

package main

import (
	"os"
	"testing"
)

var input2 string
var anchor2 int

func init() {
	data, _ := os.ReadFile("../input/day2.txt")
	input2 = string(data)
}

func BenchmarkPart2(b *testing.B) {
	var result int
	for b.Loop() {
		Answer = 0
		convertProductCodes(input2)
		result = Answer
	}
	anchor2 = result
}
